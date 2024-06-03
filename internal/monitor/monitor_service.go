package monitor

import (
	"context"
	"fmt"
	"log"

	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

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
