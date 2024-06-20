package monitor

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/go-co-op/gocron/v2"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
)

// Get up time of the server in the elasticsearch
func (m *MonitorService) GetUpTime(ctx context.Context, in *mt.UptimeRequest) (*mt.UptimeResponse, error) {
	fmt.Println("GetUpTime called in monitor service...")

	// var field = "server_ipv4"
	var duration = "duration"
	var size = 0

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
			Size:&size,
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
			time.Minute*time.Duration(m.config.MonitorDurationMinute),
		),
		gocron.NewTask(
			func(ctx context.Context) {
				go func(ctx context.Context) {
					fmt.Println("WorkDailyMonitorServer called...")

					// middleware
					mw, err := middleware.NewMiddleware(m.config.PathPublicKey)
					if err != nil {
						fmt.Println("failed to create middleware", err)
					}

					// set token to context
					token, err := mw.GetToken(m.config.TokenInternal)
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

					fmt.Println("numeber of servers: ", len(servers.Servers))

					fmt.Println("start to ping to servers...")

					// ping to each server
					m.pingToServers(ctx, servers)

					// Kênh để nhận tín hiệu hệ điều hành
					sigs := make(chan os.Signal, 1)
					signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

					// Đợi tín hiệu ngắt (Ctrl+C)
					<-sigs
					fmt.Println("Received interrupt signal, stopping workers...")
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
