package gRPCServer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/management-system/domain"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	"github.com/golang-jwt/jwt"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Ping server
func (ms *ManagementSystemGrpcServer) Ping(context.Context, *emptypb.Empty) (*managementsystem.PingResponse, error) {
	return &managementsystem.PingResponse{
		Pong: "Ping Pong",
	}, nil
}

// Create server
func (ms *ManagementSystemGrpcServer) CreateServer(ctx context.Context, in *managementsystem.CreateServerRequest) (*managementsystem.Server, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &managementsystem.Server{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	fmt.Println("token ", token)

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	fmt.Println(roles)

	if roles != "admin" {
		return &managementsystem.Server{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	// get server name, server ip and server status
	server_name := in.GetServer_Name()
	server_ipv4 := in.GetServer_IPv4()
	server_status := in.GetServer_Status()

	if server_name == "" || server_ipv4 == "" || server_status == "" {
		return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "missing server name, server ip or server status")
	}

	// check if server name already exists
	var server domain.Server
	res := ms.db.First(&server, "server_name = ?", server_name)
	// if res.Error != nil {
	// 	return &managementsystem.Server{}, status.Error(codes.Internal, "failed to query server")
	// }

	if res.RowsAffected != 0 {
		return &managementsystem.Server{}, status.Error(codes.AlreadyExists, "server name already exists")
	}

	// check if server ip already exists
	// res = ms.db.Model(&domain.Server{}).Where("server_ip = ?", server_ip).First(&server)
	res = ms.db.First(&server, "server_ipv4 = ?", server_ipv4)
	// if res.Error != nil {
	// 	return &managementsystem.Server{}, status.Error(codes.Internal, "failed to query server")
	// }

	if res.RowsAffected != 0 {
		return &managementsystem.Server{}, status.Error(codes.AlreadyExists, "server ip already exists")
	}

	// create new server
	res = ms.db.Create(&domain.Server{
		Server_Name:   server_name,
		Server_IPv4:   server_ipv4,
		Server_Status: server_status,
	})

	if res.Error != nil || res.RowsAffected == 0 {
		return &managementsystem.Server{}, status.Error(codes.Internal, "failed to create server")
	}

	var serverResponse *managementsystem.Server

	res = ms.db.Model(&domain.Server{}).Where("server_name = ?", server_name).First(&serverResponse)

	if res.Error != nil || res.RowsAffected == 0 {
		return &managementsystem.Server{}, status.Error(codes.Internal, "failed to create server")
	}

	return &managementsystem.Server{
		Server_ID:     serverResponse.Server_ID,
		Server_Name:   serverResponse.Server_Name,
		Server_IPv4:   serverResponse.Server_IPv4,
		Server_Status: serverResponse.Server_Status,
		CreatedAt:     serverResponse.CreatedAt,
		UpdatedAt:     serverResponse.UpdatedAt,
	}, nil
}

// // Update server
func (ms *ManagementSystemGrpcServer) UpdateServer(ctx context.Context, in *managementsystem.UpdateServerRequest) (*managementsystem.Server, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &managementsystem.Server{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	if roles != "admin" {
		return &managementsystem.Server{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	// get server id
	server_id := in.GetServer_ID()

	if server_id < 1 {
		return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "invalid server id")
	}

	// check if server already exists
	var server domain.Server
	res := ms.db.First(&server, "server_id = ?", server_id)

	if res.RowsAffected == 0 {
		return &managementsystem.Server{}, status.Error(codes.NotFound, "server not found")
	}

	// get server name, server ip and server status
	server_name := in.GetServer_Name()
	server_ipv4 := in.GetServer_IPv4()
	server_status := in.GetServer_Status()

	if server_name == "" && server_ipv4 == "" && server_status == "" {
		return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "missing server name, server ip or server status")
	}

	if server_name == server.Server_Name && server_ipv4 == server.Server_IPv4 && server_status == server.Server_Status {
		return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
	}

	if server_name == "" {
		if server_ipv4 == "" {
			if server_status == server.Server_Status {
				return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
			}

			// update server status
			server.Server_Status = server_status
			res = ms.db.Save(&server)

			if res.Error != nil || res.RowsAffected == 0 {
				return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
			}
		} else {
			// check if server ip already exists
			var server1 domain.Server
			res = ms.db.First(&server1, "server_ipv4 = ? AND server_id <> ?", server_ipv4, server_id)

			if res.RowsAffected != 0 {
				return &managementsystem.Server{}, status.Error(codes.AlreadyExists, "server ip already exists")
			}

			// update server ipv4
			if server.Server_IPv4 == server_ipv4 {
				if server_status == server.Server_Status || server_status == "" {
					return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
				}
			}

			if server_status != server.Server_Status && server_status != "" {
				server.Server_Status = server_status
			}

			if server.Server_IPv4 != server_ipv4 {
				server.Server_IPv4 = server_ipv4
			}

			res = ms.db.Save(&server)
			if res.Error != nil || res.RowsAffected == 0 {
				return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
			}
		}
	} else {
		// check if server name already exists
		var server1 domain.Server
		res = ms.db.First(&server1, "server_name = ? AND server_id <> ?", server_name, server_id)

		if res.RowsAffected != 0 {
			return &managementsystem.Server{}, status.Error(codes.AlreadyExists, "server name already exists")
		}

		// if server.Server_Name != server_name {
		// 	server.Server_Name = server_name
		// }

		if server_ipv4 == "" {
			if server_status == "" {
				if server.Server_Name == server_name {
					return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
				}
				// update server name
				server.Server_Name = server_name
				res = ms.db.Save(&server)

				if res.Error != nil || res.RowsAffected == 0 {
					return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
				}
			} else {
				if server.Server_Status == server_status && server.Server_Name == server_name {
					return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
				}

				if server.Server_Status != server_status {
					server.Server_Status = server_status
				}

				if server.Server_Name != server_name {
					server.Server_Name = server_name
				}

				res = ms.db.Save(&server)

				if res.Error != nil || res.RowsAffected == 0 {
					return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
				}
			}
		} else {
			// check if server ip already exists
			var server2 domain.Server
			res = ms.db.First(&server2, "server_ipv4 = ? AND server_id <> ?", server_ipv4, server_id)

			if res.RowsAffected != 0 {
				return &managementsystem.Server{}, status.Error(codes.AlreadyExists, "server ip already exists")
			}

			// server.Server_IPv4 = server_ipv4

			if server_status != "" {
				if server_name == server.Server_Name && server_ipv4 == server.Server_IPv4 && server_status == server.Server_Status {
					return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
				}

				if server_status != server.Server_Status {
					server.Server_Status = server_status
				}

				if server.Server_Name != server_name {
					server.Server_Name = server_name
				}

				if server.Server_IPv4 != server_ipv4 {
					server.Server_IPv4 = server_ipv4
				}

				res = ms.db.Save(&server)

				if res.Error != nil || res.RowsAffected == 0 {
					return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
				}
			} else {
				if server_name == server.Server_Name && server_ipv4 == server.Server_IPv4 {
					return &managementsystem.Server{}, status.Error(codes.InvalidArgument, "no changes to update")
				}

				if server.Server_Name != server_name {
					server.Server_Name = server_name
				}

				if server.Server_IPv4 != server_ipv4 {
					server.Server_IPv4 = server_ipv4
				}

				res = ms.db.Save(&server)

				if res.Error != nil || res.RowsAffected == 0 {
					return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
				}
			}
		}
	}

	var serverResponse *managementsystem.Server

	res = ms.db.Model(&domain.Server{}).Where("server_id = ?", server_id).First(&serverResponse)

	if res.Error != nil || res.RowsAffected == 0 {
		return &managementsystem.Server{}, status.Error(codes.Internal, "failed to update server")
	}

	return &managementsystem.Server{
		Server_ID:     serverResponse.Server_ID,
		Server_Name:   serverResponse.Server_Name,
		Server_IPv4:   serverResponse.Server_IPv4,
		Server_Status: serverResponse.Server_Status,
		CreatedAt:     serverResponse.CreatedAt,
		UpdatedAt:     serverResponse.UpdatedAt,
	}, nil
}

// Delete server
func (ms *ManagementSystemGrpcServer) DeleteServer(ctx context.Context, in *managementsystem.DeleteServerRequest) (*emptypb.Empty, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	if roles != "admin" {
		return &emptypb.Empty{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	// get server id
	server_id := in.GetServer_ID()

	if server_id < 1 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "invalid server id")
	}

	// check if server already exists
	var server domain.Server
	res := ms.db.First(&server, "server_id = ?", server_id)

	if res.RowsAffected == 0 {
		return &emptypb.Empty{}, status.Error(codes.NotFound, "server name not found")
	}

	// delete server
	res = ms.db.Delete(&domain.Server{}, "server_id = ?", server_id)

	if res.Error != nil || res.RowsAffected == 0 {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to delete server")
	}

	return &emptypb.Empty{}, nil
}

// Import server
func (ms *ManagementSystemGrpcServer) ImportServer(stream managementsystem.ManagementSystem_ImportServerServer) error {
	// fmt.Println("start import server")

	type server struct {
		Server_Name  string
		Server_IPv4   string
		Server_Status string
	}

	importSucces := 0
	listServersImportSucces := make([]server, 0)

	importFailed := 0
	listServersImportFailed := make([]server, 0)

	type response struct {
		ImportSucces            int
		ImportFailed            int
		ListServersImportSucces []server
		ListServersImportFailed []server
	}

	////
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			responses, _ := json.Marshal(&response{
				ImportSucces:            importSucces,
				ListServersImportSucces: listServersImportSucces,
				ImportFailed:            importFailed,
				ListServersImportFailed: listServersImportFailed,
			})

			fmt.Println("Result", importSucces, importFailed)

			return stream.SendAndClose(&httpbody.HttpBody{
				ContentType: "application/json",
				Data:        []byte(responses),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		// TODO:
		var res server
		err = json.Unmarshal(req.Content, &res)
		if err != nil {
			log.Fatalf("Error while unmarshalling: %v", err)
		}
		log.Println(res)

		// check if infomation server invalid
		if res.Server_Name == "" || res.Server_IPv4 == "" || res.Server_Status == "" {
			importFailed++
			listServersImportFailed = append(listServersImportFailed, res)
			continue
		}

		fmt.Println("server info not null")

		// check if server name already exists
		var server1 domain.Server
		result := ms.db.First(&server1, "server_name = ?", res.Server_Name)
		if result.RowsAffected != 0 {
			importFailed++
			listServersImportFailed = append(listServersImportFailed, res)
			continue
		}

		fmt.Println("server name is not existing")

		// check if server ip already exists
		var server2 domain.Server
		result = ms.db.First(&server2, "server_ipv4 = ?", res.Server_IPv4)
		if result.RowsAffected != 0 {
			importFailed++
			listServersImportFailed = append(listServersImportFailed, res)
			continue
		}

		fmt.Println("server ip is not existing")

		// create new server
		result = ms.db.Create(&domain.Server{
			Server_Name:   res.Server_Name,
			Server_IPv4:   res.Server_IPv4,
			Server_Status: res.Server_Status,
		})

		if result.RowsAffected == 0 {
			importFailed++
			listServersImportFailed = append(listServersImportFailed, res)
			continue
		}

		fmt.Println("server ip is created")

		importSucces++
		listServersImportSucces = append(listServersImportSucces, res)
	}
}
