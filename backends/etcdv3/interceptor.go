package etcdv3

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// unaryInterceptor build a UnaryClientInterceptor with custom headers.
func unaryInterceptor(headers map[string]string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(contextWithHeaders(ctx, headers), method, req, reply, cc, opts...)
	}
}

// streamInterceptor build a StreamClientInterceptor with custom headers.
func streamInterceptor(headers map[string]string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
		opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(contextWithHeaders(ctx, headers), desc, cc, method, opts...)
	}
}

// contextWithHeaders append headers to OutgoingContext.
func contextWithHeaders(ctx context.Context, headers map[string]string) context.Context {
	for k, v := range headers {
		ctx = metadata.AppendToOutgoingContext(ctx, k, v)
	}
	return ctx
}
