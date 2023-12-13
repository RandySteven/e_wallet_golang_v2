// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/e_wallet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	RegisterUser(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserResponse, error)
	LoginUser(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordToken, error)
	ResetPassword(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/e_wallet.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginUser(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/e_wallet.UserService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordToken, error) {
	out := new(ForgotPasswordToken)
	err := c.cc.Invoke(ctx, "/e_wallet.UserService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/e_wallet.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	RegisterUser(context.Context, *UserRegisterRequest) (*UserResponse, error)
	LoginUser(context.Context, *UserLoginRequest) (*UserResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordToken, error)
	ResetPassword(context.Context, *PasswordResetRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) RegisterUser(context.Context, *UserRegisterRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) LoginUser(context.Context, *UserLoginRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *PasswordResetRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.UserService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginUser(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.UserService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*PasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "e_wallet.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserService_LoginUser_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _UserService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/e_wallet.proto",
}

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	CreateTransferTransaction(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransactionDetailResponse, error)
	CreateTopupTransaction(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TransactionDetailResponse, error)
	GetAllTransactionsRecords(ctx context.Context, in *QueryCondition, opts ...grpc.CallOption) (*TransactionPaginationResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) CreateTransferTransaction(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransactionDetailResponse, error) {
	out := new(TransactionDetailResponse)
	err := c.cc.Invoke(ctx, "/e_wallet.TransactionService/CreateTransferTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateTopupTransaction(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TransactionDetailResponse, error) {
	out := new(TransactionDetailResponse)
	err := c.cc.Invoke(ctx, "/e_wallet.TransactionService/CreateTopupTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetAllTransactionsRecords(ctx context.Context, in *QueryCondition, opts ...grpc.CallOption) (*TransactionPaginationResponse, error) {
	out := new(TransactionPaginationResponse)
	err := c.cc.Invoke(ctx, "/e_wallet.TransactionService/GetAllTransactionsRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	CreateTransferTransaction(context.Context, *TransferRequest) (*TransactionDetailResponse, error)
	CreateTopupTransaction(context.Context, *TopupRequest) (*TransactionDetailResponse, error)
	GetAllTransactionsRecords(context.Context, *QueryCondition) (*TransactionPaginationResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) CreateTransferTransaction(context.Context, *TransferRequest) (*TransactionDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransferTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) CreateTopupTransaction(context.Context, *TopupRequest) (*TransactionDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopupTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) GetAllTransactionsRecords(context.Context, *QueryCondition) (*TransactionPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTransactionsRecords not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_CreateTransferTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransferTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.TransactionService/CreateTransferTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransferTransaction(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateTopupTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTopupTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.TransactionService/CreateTopupTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTopupTransaction(ctx, req.(*TopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetAllTransactionsRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCondition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetAllTransactionsRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e_wallet.TransactionService/GetAllTransactionsRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetAllTransactionsRecords(ctx, req.(*QueryCondition))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "e_wallet.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransferTransaction",
			Handler:    _TransactionService_CreateTransferTransaction_Handler,
		},
		{
			MethodName: "CreateTopupTransaction",
			Handler:    _TransactionService_CreateTopupTransaction_Handler,
		},
		{
			MethodName: "GetAllTransactionsRecords",
			Handler:    _TransactionService_GetAllTransactionsRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/e_wallet.proto",
}
