package monitor

import (
	"context"
	"encoding/json"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/segmentio/kafka-go"
)

func (m *MonitorService) StartMonitorConsumer(ctx context.Context) {
	for {
		msg, err := m.MonitorConsumer.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		go m.Worker(msg)
	}
}

// Indexing a document
// The standard way of indexing a document is to provide a struct to the Request method,
// the standard json/encoder will be run on your structure
// and the result will be sent to Elasticsearch.
type Document struct {
	Timestamp   time.Time `json:"timestamp"`
	Server_IPv4 string    `json:"server_ipv4"`
	Duration    int       `json:"duration"`
}

func (m *MonitorService) Worker(msg kafka.Message) {
	// fmt.Println(string(msg.Topic))

	// result topic from management system
	var results struct {
		TimeMonitor time.Time `json:"time_monitor"`
		ServerIPv4s []string  `json:"server_ipv4s"`
	}

	// Unmarshal JSON to the struct/map for efficient message decoding
	err := json.Unmarshal(msg.Value, &results)
	if err != nil {
		m.logger.Error("Failed to unmarshal server: " + err.Error())
		return // Handle error appropriately
	}

	// fmt.Println(results)

	// {
	// 	timestamp: 9:55
	// 	server_id: 1
	// 	duration: 10
	// }

	// _timestamp, _ := time.Parse("2006-01-02 15:04:05", results.TimeMonitor)
	// fmt.Println(_timestamp)
	for _, serverIPv4 := range results.ServerIPv4s {
		_, _ = m.elasticClient.Index("uptime-server-monitor").
			Document(&Document{
				Timestamp:   results.TimeMonitor,
				Server_IPv4: serverIPv4,
				Duration:    10,
			}).
			Refresh(refresh.Waitfor).
			Do(context.Background())
	}
}
