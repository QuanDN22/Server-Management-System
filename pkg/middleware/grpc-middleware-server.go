package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (mw *Middleware) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println()
	fmt.Println("4. Unary Server Interceptor")
	fmt.Println(info.FullMethod)
	if info.FullMethod == "/auth.service.AuthService/Login" || info.FullMethod == "/auth.service.AuthService/Signup" {
		fmt.Println("4. start /auth.service.AuthService/Login")
		return handler(ctx, req)
	}

	fmt.Println("4. start")

	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
	}

	tokens := headers.Get("jwt")
	if len(tokens) < 1 {
		return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
	}

	tokenString := tokens[0] // just use the first, ignore repeated headers

	token, err := mw.GetToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	ctx = ContextSetToken(ctx, token)
	return handler(ctx, req)
}
