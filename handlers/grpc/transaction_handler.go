package handler_grpc

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/proto"
	"github.com/shopspring/decimal"
)

type TransactionHandler struct {
	pb.UnimplementedTransactionServiceServer
	usecase interfaces.TransactionUsecase
}

func NewTransactionHandler(usecase interfaces.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{usecase: usecase}
}

func (h *TransactionHandler) CreateTransferTransaction(ctx context.Context, request *pb.TransferRequest) (*pb.TransactionDetailResponse, error) {
	userId := ctx.Value("id").(uint)
	req := &req.TransferRequest{
		SenderUserId:     userId,
		ReceiverWalletId: request.ReceiverWalletId,
		Amount:           decimal.NewFromFloat32(request.Amount),
		Description:      request.Description,
	}
	res, err := h.usecase.CreateTransferTransaction(ctx, req)
	if err != nil {
		return nil, err
	}
	amount, _ := res.Amount.Float64()
	resp := &pb.TransactionDetailResponse{
		Id:              uint32(res.ID),
		TransactionDate: res.CreatedAt.Format("2006-01-02"),
		// SenderName:      res.Sender.User.Name,
		// SenderWallet:    res.Sender.Number,
		// ReceipentName:   res.Receiver.User.Name,
		// ReceipentWallet: res.Receiver.Number,
		Description: res.Description,
		Amount:      float32(amount),
	}
	return resp, nil
}

func (h *TransactionHandler) CreateTopupTransaction(ctx context.Context, request *pb.TopupRequest) (*pb.TransactionDetailResponse, error) {
	userId := ctx.Value("id").(uint)
	req := &req.TopupRequest{
		UserID:       userId,
		Amount:       decimal.NewFromFloat32(request.Amount),
		SourceOfFund: uint(request.SourceOfFundId),
	}
	res, err := h.usecase.CreateTopupTransaction(ctx, req)
	if err != nil {
		return nil, err
	}
	amount, _ := res.Amount.Float64()
	resp := &pb.TransactionDetailResponse{
		ReceipentName:   res.Receiver.User.Name,
		ReceipentWallet: res.Receiver.Number,
		Description:     res.Description,
		Amount:          float32(amount),
	}
	return resp, nil
}

func (h *TransactionHandler) GetAllTransactionsRecords(ctx context.Context, request *pb.QueryCondition) (*pb.TransactionPaginationResponse, error) {
	return nil, nil
}
