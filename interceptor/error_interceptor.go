package interceptor

import (
	"context"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apperror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	var (
		errFieldValidation       *apperror.ErrFieldValidation
		errAmountLimit           *apperror.ErrAmountLimit
		errBalanceNotEnough      *apperror.ErrBalanceNotEnough
		errSenderAndReceiverSame *apperror.ErrSenderAndReceiverSame
		errDataNotFound          *apperror.ErrDataNotFound
		errWalletInvalid         *apperror.ErrWalletInvalid
		errInvalidRequest        *apperror.ErrInvalidRequest
		errTokenAlreadyUsed      *apperror.ErrTokenAlreadyUsed
		errTokenExpired          *apperror.ErrTokenExpired
		errTokenInvalid          *apperror.ErrTokenInvalid
		errInvalidFormat         *apperror.ErrInvalidFormat
		errLogin                 *apperror.ErrLogin
		errEmailAlreadyExists    *apperror.ErrEmailAlreadyExists
	)

	if err != nil {
		switch {
		case errors.As(err, &errSenderAndReceiverSame):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errAmountLimit):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errWalletInvalid):
		case errors.As(err, &errDataNotFound):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errBalanceNotEnough):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errInvalidRequest):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errTokenAlreadyUsed):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errTokenExpired):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errTokenInvalid):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errFieldValidation):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errInvalidFormat):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errLogin):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errEmailAlreadyExists):
			return res, status.Error(codes.AlreadyExists, err.Error())
		default:
			return res, status.Error(codes.Unknown, err.Error())
		}
	}

	return res, err
}
