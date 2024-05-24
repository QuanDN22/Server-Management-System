package monitor

import (
	"context"

	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
	"github.com/golang/protobuf/ptypes/empty"
)

func (m *MonitorService) GetUpTime(ctx context.Context, _ *empty.Empty) (*mt.ResponseUptime, error) {
	// Get uptime of the server

	// return uptime
	return &mt.ResponseUptime{
		Uptime: 10.00,
	}, nil
}
