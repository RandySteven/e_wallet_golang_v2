package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// md, ok := metadata.FromIncomingContext(ctx)

	if info.FullMethod == "" {
		return handler(ctx, req)
	}

	// if !ok {
	// 	return nil, status.Error(codes.Unauthenticated, "unauthroized")
	// }

	// tokenHeader := md["authorization"][0]
	// token := utils.ValidateToken(tokenHeader)
	// if token == nil {
	// 	return nil, status.Error(codes.Unauthenticated, "invalid token")
	// }

	// ctx = context.WithValue(ctx, "id", token.ID)
	// ctx = context.WithValue(ctx, "name", token.Name)
	// ctx = context.WithValue(ctx, "email", token.Email)

	return handler(ctx, req)
}

func ErrorInterceptor() {}

func LoggerInterceptor() {}
