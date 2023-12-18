package interceptor

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if info.FullMethod == "/e_wallet.UserService/LoginUser" || info.FullMethod == "/e_wallet.UserService/RegisterUser" {
		return handler(ctx, req)
	}

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthroized")
	}

	tokenHeader := md["authorization"][0]
	token := utils.ValidateToken(tokenHeader)
	if token == nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	ctx = context.WithValue(ctx, "id", token.ID)
	ctx = context.WithValue(ctx, "name", token.Name)
	ctx = context.WithValue(ctx, "email", token.Email)

	return handler(ctx, req)
}
