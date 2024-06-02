package gRPCServer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/QuanDN22/Server-Management-System/internal/management-system/domain"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	"github.com/golang-jwt/jwt"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/QuanDN22/Server-Management-System/proto/mail"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
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

// Update server
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
		Server_Name   string
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
			responses, _ := json.MarshalIndent(&response{
				ImportSucces:            importSucces,
				ListServersImportSucces: listServersImportSucces,
				ImportFailed:            importFailed,
				ListServersImportFailed: listServersImportFailed,
			}, "", " ")

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

// View server
func (ms *ManagementSystemGrpcServer) ViewServer(ctx context.Context, in *managementsystem.ViewServerRequest) (*managementsystem.ViewServerResponse, error) {
	// get data in redis
	// key, _ := json.Marshal(in)

	// data, err := ms.cache.Get(ctx, string(key)).Result()

	// if err == nil {
	// 	return &managementsystem.ViewServerResponse{
	// 		Content: []byte(data),
	// 	}, nil
	// }

	// if not found in redis, get data in database
	limit := in.GetLimit()
	offset := in.GetOffset()
	filter_server_name := in.GetFilterServerName()
	filter_server_ipv4 := in.GetFilterServerIpv4()
	filter_server_status := in.GetFilterServerStatus()
	sort := in.GetSort()

	fmt.Println(limit, offset, filter_server_name, filter_server_ipv4, filter_server_status, sort)

	var servers []domain.Server

	// Fetch data from database
	result := ms.db.Model(&domain.Server{})

	// Apply filters
	fmt.Printf("server_name: %s\n", filter_server_name)
	if filter_server_name != "" {
		result = ms.db.Where("server_name LIKE ?", "%"+filter_server_name+"%")
	}
	fmt.Printf("number 1: %d", result.RowsAffected)

	fmt.Printf("server_ipv4: %s\n", filter_server_ipv4)
	if filter_server_ipv4 != "" {
		result = result.Where("server_ipv4 LIKE ?", "%"+filter_server_ipv4+"%")
	}
	fmt.Printf("number 2: %d\n", result.RowsAffected)

	fmt.Printf("server_status: %s\n", filter_server_status)
	if filter_server_status != "" {
		result = result.Where("server_status = ?", filter_server_status)
	}
	fmt.Printf("number 3: %d\n", result.RowsAffected)

	fmt.Printf("sort: %s\n", sort)
	if sort != "" {
		ops := strings.Split(sort, ",")
		for _, v := range ops {
			fmt.Printf("ops: %s", v)
			op := strings.Split(v, ".")
			result = result.Order(fmt.Sprintf("%s %s", op[0], op[1]))
		}
	}
	fmt.Printf("number 4: %d\n", result.RowsAffected)

	_limit, _ := strconv.Atoi(limit)
	_offset, _ := strconv.Atoi(offset)

	if _offset > 0 {
		_offset = _offset*_limit - 1
	}

	fmt.Printf("limit: %d, offset: %d\n", _limit, _offset)
	result = result.Limit(_limit).Offset(_offset)

	// Execute query
	if err := result.Find(&servers).Error; err != nil {
		// Handle error
		return nil, err
	}

	type response struct {
		Total   int
		Servers []domain.Server
	}

	res, _ := json.MarshalIndent(&response{
		Total:   int(result.RowsAffected),
		Servers: servers,
	}, "", " ")

	fmt.Println(result.RowsAffected)

	// set data in redis
	// _ = ms.cache.Set(ctx, string(key), res, 0).Err()

	return &managementsystem.ViewServerResponse{
		Content: res,
	}, nil
}

// Report server
func (ms *ManagementSystemGrpcServer) Report(ctx context.Context, in *managementsystem.ReportRequest) (*emptypb.Empty, error) {
	fmt.Println("Report server in Management System begin")

	// count server
	var count_server int64
	if err := ms.db.Model(&domain.Server{}).Count(&count_server).Error; err != nil {
		fmt.Println("Error counting records:", err)
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Error counting records")
	}

	// count server on
	var count_server_on int64
	if err := ms.db.Model(&domain.Server{}).Where("server_status LIKE ?", "on").Count(&count_server_on).Error; err != nil {
		fmt.Println("Error counting records:", err)
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Error counting records")
	}

	// count server off
	count_server_off := count_server - count_server_on

	fmt.Println(count_server, count_server_on, count_server_off)

	fmt.Println("Email:", in.GetEmail())

	// get uptime of server
	uptime, err := ms.monitorClient.GetUpTime(ctx, &mt.UptimeRequest{
		Start: in.GetStart(),
		End:   in.GetEnd(),
	})

	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	fmt.Println("uptime: ", uptime)

	type data_server struct {
		Sum_Server     int64   `json:"sum_server"`
		Sum_Server_On  int64   `json:"sum_server_on"`
		Sum_Server_Off int64   `json:"sum_server_off"`
		Uptime         float32 `json:"uptime"`
	}

	type data struct {
		Email      []string      `json:"email"`
		DataServer []data_server `json:"data_send"`
	}

	send_data, _ := json.Marshal(data{
		Email: in.GetEmail(),
		DataServer: []data_server{
			{
				Sum_Server:     count_server,
				Sum_Server_On:  count_server_on,
				Sum_Server_Off: count_server_off,
				Uptime:         uptime.Uptime,
			},
		},
	})

	fmt.Println("prepare send mail")

	// send mail
	_, err = ms.mailClient.SendMail(ctx, &mail.SendMailRequest{
		DataSendMail: send_data,
	})

	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
