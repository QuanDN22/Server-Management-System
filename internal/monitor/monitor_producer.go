package monitor

import (
	"context"
	"encoding/binary"
	"time"

	"github.com/segmentio/kafka-go"
)

func (m *MonitorService) StartMonitorProducer(ctx context.Context) {
	for {
		time_monitor := time.Now().Unix()
		time_byte := make([]byte, 8)
		binary.LittleEndian.PutUint64(time_byte, uint64(time_monitor))
		err := m.MonitorProducer.WriteMessages(ctx, kafka.Message{
			Key:   []byte("monitor"),
			Value: []byte(time_byte),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		time.Sleep(time.Second / 1000)
	}
}


