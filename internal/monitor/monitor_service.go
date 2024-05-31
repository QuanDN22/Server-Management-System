package monitor

import (
	"context"
	"fmt"
	"log"

	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/golang/protobuf/ptypes/empty"
)

func (m *MonitorService) GetUpTime(ctx context.Context, _ *empty.Empty) (*mt.ResponseUptime, error) {
	var field = "server_id"
	var duration = "duration"
	var from = "2024-05-29T18:30:00+07:00"
	var to = "2024-05-30T18:32:00+07:00"

	// Get uptime of the server
	result, err := m.elasticClient.Search().Index("uptime-server-monitor").
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Filter: []types.Query{
						types.Query{
							Range: map[string]types.RangeQuery{
								"timestamp": types.DateRangeQuery{
									Gte: &from,
									Lte: &to,
								},
							},
						},
					},
				},
			},
			Aggregations: map[string]types.Aggregations{
				"servers": types.Aggregations{
					Terms: &types.TermsAggregation{
						Field: &field,
					},
					Aggregations: map[string]types.Aggregations{
						"total_duration": types.Aggregations{
							Sum: &types.SumAggregation{
								Field: &duration,
							},
						},
					},
				},
			},
		}).Do(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	// Extract the sum from the aggregation result
	fmt.Printf("Total uptime: %v", result.Aggregations["total_duration"].(*types.SimpleValueAggregate).Value)

	// return uptime
	return &mt.ResponseUptime{
		Uptime: float32(result.Aggregations["total_duration"].(*types.SimpleValueAggregate).Value),
	}, nil
}
