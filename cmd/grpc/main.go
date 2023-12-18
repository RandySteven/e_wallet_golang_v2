package main

import (
	"log"
	"net"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apps"
	handler_grpc "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/handlers/grpc"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interceptor"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		return
	}

	opt := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor.AuthInterceptor, interceptor.ErrorInterceptor, interceptor.LoggerInterceptor),
	}

	server := grpc.NewServer(opt...)

	repository := InitRepository()
	usecase := apps.NewUsecase(*repository)

	userHandler := handler_grpc.NewUserHandler(usecase.UserUsecase)
	transactionHandler := handler_grpc.NewTransactionHandler(usecase.TransactionUsecase)

	pb.RegisterUserServiceServer(server, userHandler)
	pb.RegisterTransactionServiceServer(server, transactionHandler)

	if err = server.Serve(listener); err != nil {
		return
	}

}
