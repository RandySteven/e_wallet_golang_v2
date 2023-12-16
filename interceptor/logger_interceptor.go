package interceptor

import (
	"context"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	log := logger.NewLog()

	start := time.Now()
	args := map[string]interface{}{
		"status":  status.Code(err),
		"latency": time.Since(start),
		"path":    info.FullMethod,
	}

	if err != nil {
		args["err"] = err.Error()
		log.Error(args)
	}

	return res, err
}
