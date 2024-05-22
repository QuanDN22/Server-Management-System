package monitor

import (
	"context"
	"encoding/json"
	"fmt"

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

func (m *MonitorService) Worker(msg kafka.Message) {
	fmt.Println(string(msg.Topic))

	// result topic from management system
	var result struct {
		TimeMonitor string `json:"time_monitor"`
		ServerIDs   []uint `json:"server_ids"`
	}

	// Unmarshal JSON to the struct/map for efficient message decoding
	err := json.Unmarshal(msg.Value, &result)
	if err != nil {
		m.logger.Error("Failed to unmarshal server: " + err.Error())
		return // Handle error appropriately
	}

	fmt.Println(result)
}
