package handler_grpc

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase interfaces.UserUsecase
}

// ForgotPassword implements proto.UserServiceServer.
func (h *UserHandler) ForgotPassword(ctx context.Context, request *pb.ForgotPasswordRequest) (*pb.ForgotPasswordToken, error) {
	panic("unimplemented")
}

// LoginUser implements proto.UserServiceServer.
func (h *UserHandler) LoginUser(ctx context.Context, request *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	loginReq := &req.UserLoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.usecase.LoginUser(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	return &pb.UserLoginResponse{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}, nil
}

// RegisterUser implements proto.UserServiceServer.
func (h *UserHandler) RegisterUser(ctx context.Context, request *pb.UserRegisterRequest) (*pb.UserResponse, error) {

	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.usecase.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// ResetPassword implements proto.UserServiceServer.
func (h *UserHandler) ResetPassword(ctx context.Context, request *pb.PasswordResetRequest) (*pb.UserResponse, error) {
	panic("unimplemented")
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ pb.UserServiceServer = &UserHandler{}
