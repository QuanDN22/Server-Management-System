package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (mw *Middleware) UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if method == "/auth.service.AuthService/Login" || method == "/auth.service.AuthService/Signup" {
		fmt.Println("3. start /auth.service.AuthService/Login")
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	fmt.Println("3. start")

	token, err := ContextGetToken(ctx)
	if err != nil {
		return fmt.Errorf("token not set in context: %w", err)
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "jwt", token.Raw)

	// call the invoker with everythign else untouched
	return invoker(ctx, method, req, reply, cc, opts...)
}
