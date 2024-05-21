package gRPCServer

import (
	"context"

	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Ping server
func (ms *ManagementSystemGrpcServer) Ping(context.Context, *emptypb.Empty) (*managementsystem.PingResponse, error) {
	return &managementsystem.PingResponse{
		Pong: "Ping Pong",
	}, nil
}

// // Create server
// func (ms *ManagementSystemGrpcServer) CreateServer(context.Context, *managementsystem.CreateServerRequest) (*managementsystem.Server, error)
// // Update server
// func (ms *ManagementSystemGrpcServer) UpdateServer(context.Context, *managementsystem.UpdateServerRequest) (*managementsystem.Server, error)
// // Delete server
// func (ms *ManagementSystemGrpcServer) DeleteServer(context.Context, *managementsystem.DeleteServerRequest) (*emptypb.Empty, error)
