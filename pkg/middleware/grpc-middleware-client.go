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

func (mw *Middleware) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("3. Stream Client Interceptor")
	fmt.Println(method)

	token, err := ContextGetToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("token not set in context: %w", err)
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "jwt", token.Raw)

	// Create a wrapper stream to carry the modified context
	// wrappedStreamer := &wrapperStreamer{Streamer: streamer, ctx: ctx}

	// Call the actual streamer with the modified context
	return streamer(ctx, desc, cc, method, opts...)
}
