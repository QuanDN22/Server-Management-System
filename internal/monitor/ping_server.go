package monitor

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	management_system "github.com/QuanDN22/Server-Management-System/proto/management-system"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/segmentio/kafka-go"
)

func (m *MonitorService) pingToServers(ctx context.Context, servers *management_system.GetAllServerResponse) {
	// wait group
	wg := sync.WaitGroup{}

	// create a buffer channel
	buffer := make(chan int, m.config.MaxConurrentPingServers)

	// ping to each server_ipv4
	// if server is on, save in the elasticsearch
	// send data to the management system via kafka to update the status of the server in the database
	for _, server := range servers.Servers {
		fmt.Println("server infomation: ", server.GetServer_ID(), server.GetServer_IPv4())

		// add goroutine to wait group
		wg.Add(1)

		// push to buffer
		buffer <- 1

		// ping to server ipv4
		go m.executePingToServer(ctx, server.GetServer_ID(), server.GetServer_IPv4(), &wg, buffer)
	}

	// wait for the buffer to be empty
	wg.Wait()
}

// execute ping to server
func (m *MonitorService) executePingToServer(ctx context.Context, server_id int64, server_ipv4 string, wg *sync.WaitGroup, buffer chan int) {
	// defer wait group done
	defer wg.Done()

	// ping to server ipv4
	out, err := exec.Command("ping", server_ipv4).Output()

	if err != nil {
		fmt.Printf("Error in pinging to server, %v", err)
		<-buffer
		return
	}

	server_status := "off"

	// if server is on, save in the elasticsearch
	if strings.Contains(string(out), "bytes=") {
		fmt.Println(string(out))

		server_status = "on"

		// save in the elasticsearch
		type Document struct {
			Timestamp   time.Time `json:"timestamp"`
			Server_ID   int64     `json:"server_id"`
			Server_IPv4 string    `json:"server_ipv4"`
			Duration    int       `json:"duration"`
		}

		_, _ = m.elasticClient.Index("uptime-server-monitor").
			Document(&Document{
				Timestamp:   time.Now(),
				Server_ID:   server_id,
				Server_IPv4: server_ipv4,
				Duration:    m.config.MonitorDurationMinute,
			}).
			Refresh(refresh.Waitfor).
			Do(context.Background())

		// send data to the management system to update the status of the server in the database
		err = m.MonitorProducer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(fmt.Sprint(server_id)),
			Value: []byte(fmt.Sprintf("%d,%s", server_id, server_status)),
		})
		if err != nil {
			<-buffer
			fmt.Println("could not write message " + err.Error())
			return
		}
	}

	// pop from buffer
	<-buffer
}
