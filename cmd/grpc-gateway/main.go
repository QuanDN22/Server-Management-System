package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "github.com/QuanDN22/Server-Management-System/proto/auth"
	mspb "github.com/QuanDN22/Server-Management-System/proto/management-system"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating gRPC-gateway")

	// config
	cfg, err := config.NewConfig("./cmd/grpc-gateway", ".env.grpc-gateway")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

	// new logger
	// Create a logger with lumberjack integration
	l, err := logger.NewLogger(
		fmt.Sprintf("%s%s.log", cfg.LogFilename, cfg.ServiceName),
		int(cfg.LogMaxSize),
		int(cfg.LogMaxBackups),
		int(cfg.LogMaxAge),
		true,
		zapcore.InfoLevel,
	)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	l.Info("logger created...")

	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	// mw, err := middleware.NewMiddleware(os.Args[1])
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// Create a client connection to the gRPC server we just created
	// This is where the gRPC-gateway proxies the requests
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(mw.StreamClientInterceptor),
	}

	err = authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, cfg.AuthServerPort, opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = mspb.RegisterManagementSystemHandlerFromEndpoint(ctx, gwmux, cfg.ManagementSystemServerPort, opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// import server
	// Attachment upload from http/s handled manually
	gwmux.HandlePath("POST", "/v1/api/servers/import", handleImportServerFile)

	// export server
	// "/v1/api/servers/export"

	gwServer := &http.Server{
		Addr:    cfg.GrpcGatewayPort,
		Handler: mw.HandleHTTP(gwmux),
	}

	l.Info(fmt.Sprintf("Serving gRPC-Gateway is running on %s", cfg.GrpcGatewayPort))
	log.Fatalln(gwServer.ListenAndServe())
}

func handleImportServerFile(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// config
	cfg, err := config.NewConfig("./cmd/grpc-gateway", ".env.grpc-gateway")
	if err != nil {
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	// mw, err := middleware.NewMiddleware(os.Args[1])
	if err != nil {
		log.Println("failed to create middleware", zap.Error(err))
	}

	// validate token
	parts := strings.Split(r.Header.Get("Authorization"), " ")
	if len(parts) < 2 || parts[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("missing or invalid authorization header")) //nolint
		return
	}
	tokenString := parts[1]

	token, err := mw.GetToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid token: " + err.Error())) //nolint
		return
	}

	ctx := middleware.ContextSetToken(r.Context(), token)

	// parse form take file
	err = r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("attachment")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println(header.Filename, header.Size, header.Header)

	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.

	// read excel document
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Println(err)
		return
	}

	// remove first element in the slice rows
	rows = append(rows[:0], rows[1:]...)

	///////

	//create a connection
	conn, err := grpc.Dial(
		cfg.ManagementSystemServerPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStreamInterceptor(mw.StreamClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//create a client
	client := mspb.NewManagementSystemClient(conn)

	stream, err := client.ImportServer(ctx)

	if err != nil {
		log.Fatalf("Error while calling Upload: %v", err)
	}

	type server struct {
		Server_Name  string
		Server_IPv4   string
		Server_Status string
	}

	// add server in database with three 3 fields: server_name, server_ip, server_status
	for _, row := range rows {
		data, _ := json.Marshal(&server{
			Server_Name:  row[0],
			Server_IPv4:   row[1],
			Server_Status: row[2],
		})

		stream.Send(&mspb.ImportServerRequest{
			Content: data,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from ImportServerFile: %v\n", err)
	}

	w.Write(res.Data)
}
