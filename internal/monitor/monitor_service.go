package monitor

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/go-co-op/gocron/v2"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
)

// Get up time of the server in the elasticsearch
func (m *MonitorService) GetUpTime(ctx context.Context, in *mt.UptimeRequest) (*mt.UptimeResponse, error) {
	fmt.Println("GetUpTime called in monitor service...")

	// var field = "server_ipv4"
	var duration = "duration"

	start_ := in.GetStart() + "+07:00"
	end_ := in.GetEnd() + "+07:00"

	fmt.Println("monitor service getuptime: ", start_, " ", end_)

	// Get uptime of the server
	result, err := m.elasticClient.Search().Index("uptime-server-monitor").
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Filter: []types.Query{
						{
							Range: map[string]types.RangeQuery{
								"timestamp": types.DateRangeQuery{
									Gte: &start_,
									Lte: &end_,
								},
							},
						},
					},
				},
			},
			Aggregations: map[string]types.Aggregations{
				"total_duration": {
					Sum: &types.SumAggregation{
						Field: &duration,
					},
				},
			},
		}).Do(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	// Extract the sum from the aggregation result
	fmt.Printf("Total uptime: %v", result.Aggregations["total_duration"].(*types.SumAggregate).Value)

	// return uptime
	return &mt.UptimeResponse{
		Uptime: float32(result.Aggregations["total_duration"].(*types.SumAggregate).Value),
	}, nil
}

// Aggregations: map[string]types.Aggregations{
// 	"servers": types.Aggregations{
// 		Terms: &types.TermsAggregation{
// 			Field: &field,
// 		},
// 		Aggregations: map[string]types.Aggregations{
// 			"total_duration": types.Aggregations{
// 				Sum: &types.SumAggregation{
// 					Field: &duration,
// 				},
// 			},
// 		},
// 	},
// },

// monitor server
func (m *MonitorService) WorkDailyMonitorServer(ctx context.Context) {
	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
		fmt.Println("Error in creating scheduler")
	}

	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.DurationJob(
			// time.Minute * time.Duration(m.config.MonitorDurationMinute),
			time.Second*10,
		),
		gocron.NewTask(
			func(ctx context.Context) {
				go func(ctx context.Context) {
					fmt.Println("WorkDailyMonitorServer called...")

					// // middleware
					// mw, err := middleware.NewMiddleware(m.config.PathPublicKey)
					// if err != nil {
					// 	fmt.Println("failed to create middleware", err)
					// }

					// // set token to context
					// token, err := mw.GetToken(m.config.TokenInternal)
					// if err != nil {
					// 	fmt.Println(("invalid token: " + err.Error()))
					// 	return
					// }

					// login to take token
					login, err := m.authClient.Login(ctx, &auth.LoginRequest{
						Username: "admin2",
						Password: "2",
					})

					if err != nil {
						fmt.Println("Error login:", err)
						return
					}

					fmt.Println("Token:", login.AccessToken)

					// middleware
					mw, err := middleware.NewMiddleware(m.config.PathPublicKey)
					if err != nil {
						fmt.Println("failed to create middleware", err)
					}

					// set token to context
					token, err := mw.GetToken(login.AccessToken)
					if err != nil {
						fmt.Println(("invalid token: " + err.Error()))
						return
					}

					ctx = middleware.ContextSetToken(ctx, token)

					// get all server including server_id and server_ipv4
					servers, err := m.managementClient.GetAllServer(ctx, &emptypb.Empty{})
					if err != nil {
						fmt.Printf("Error in getting all server %v", err)
					}

					// ping to each server_ipv4
					// if server is on, save in the elasticsearch
					// send data to the management system via kafka to update the status of the server in the database
					for _, server := range servers.GetServers() {
						fmt.Println("server infomation: ", server.GetServer_IPv4(), server.GetServer_ID())

						// ping to server ipv4
						ping_result, err := pingToServer(server.GetServer_IPv4())

						if err != nil {
							fmt.Println("Error in pinging to server")
						}

						server_status := "off"

						// if server is on, save in the elasticsearch
						if ping_result {
							server_status = "on"

							type Document struct {
								Timestamp   time.Time `json:"timestamp"`
								Server_ID   int64     `json:"server_id"`
								Server_IPv4 string    `json:"server_ipv4"`
								Duration    int       `json:"duration"`
							}

							fmt.Println(int(m.config.MonitorDurationMinute))

							_, _ = m.elasticClient.Index("uptime-server-monitor").
								Document(&Document{
									Timestamp:   time.Now(),
									Server_ID:   server.GetServer_ID(),
									Server_IPv4: server.GetServer_IPv4(),
									Duration:    m.config.MonitorDurationMinute, 
								}).
								Refresh(refresh.Waitfor).
								Do(context.Background())
						}

						// send data to the management system to update the status of the server in the database
						err = m.MonitorProducer.WriteMessages(ctx, kafka.Message{
							Key:   []byte(fmt.Sprint(server.GetServer_ID())),
							Value: []byte(fmt.Sprintf("%d,%s", server.GetServer_ID(), server_status)),
						})
						if err != nil {
							panic("could not write message " + err.Error())
						}
					}
				}(ctx)
			},
			ctx,
		),
	)
	if err != nil {
		// handle error
		fmt.Println("Error in adding job to scheduler")
	}

	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	c := make(chan byte)
	<-c

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
		fmt.Println("Error in shutting down scheduler")
	}
}
