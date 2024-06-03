package ping

import (
	"context"
	"fmt"
	"time"

	"github.com/go-ping/ping"
	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
)

type PingService struct {
	logger       *zap.Logger
	PingProducer *kafka.Writer

	manageClient managementsystem.ManagementSystemClient
}

// PingBrokerAddress string, PingTopic string,

func NewPingService(
	ctx context.Context,
	PingProducer *kafka.Writer,
	logger *zap.Logger,
	manageClient managementsystem.ManagementSystemClient,
) *PingService {
	return &PingService{
		PingProducer: PingProducer,
		logger:       logger,
		manageClient: manageClient,
	}
}

func (p *PingService) Start(ctx context.Context, numberOfServer uint) {
	time.Sleep(time.Second * 3)
	for {
		// get all server_ipv4
		server_ipv4s, err := p.manageClient.GetAllServerIP(ctx, &emptypb.Empty{})
		if err != nil {
			continue
		}

		for _, server_ipv4 := range server_ipv4s.Server_IPv4 {
			// ping to server
			pinger, err := ping.NewPinger(server_ipv4)
			if err != nil {
				panic(err)
			}

			pinger.Count = 3
			pinger.Run()                 // blocks until finished
			stats := pinger.Statistics() // get send/receive/rtt stats

			var status string
			if stats.PacketLoss != 3 {
				status = "on"
			}

			// send to kafka
			err = p.PingProducer.WriteMessages(ctx, kafka.Message{
				Key:   []byte(server_ipv4),
				Value: []byte(fmt.Sprintf("%s,%s", server_ipv4, status)),
			})
			if err != nil {
				panic("could not write message " + err.Error())
			}

			msg := fmt.Sprintf(`{"server_id": %s, "status": %s}`, server_ipv4, status)
			p.logger.Info(msg)
		}

		time.Sleep(time.Second * 1)
	}
}
