package handlers

import (
	"assignment_4/apperror"
	"assignment_4/entities"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"assignment_4/enums"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	uscase interfaces.TransactionUsecase
}

// GetAllTransactionsRecords implements interfaces.TransactionHandler.
func (handler *TransactionHandler) GetAllTransactionsRecords(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
	)

	var query entities.QueryCondition
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Error(err)
		return
	}

	transactions, err := handler.uscase.GetAllTransactionsRecords(ctx, &query)
	if err != nil {
		c.Error(err)
		return
	}

	var transactionDetails []res.TransactionDetailResponse
	for _, transaction := range transactions {
		transactionDetail := res.TransactionDetailResponse{
			ID:              transaction.ID,
			TransactionDate: transaction.CreatedAt,
			Description:     transaction.Description,
			Amount:          transaction.Amount,
		}

		if transaction.ReceiverID == transaction.SenderID {
			transactionDetail.TransactionType = enums.Topup
			transactionDetail.TopupUser = transaction.Sender.User.Name
			transactionDetail.TopupWallet = transaction.Sender.Number
		} else {
			transactionDetail.TransactionType = enums.Transfer
			transactionDetail.SenderName = transaction.Sender.User.Name
			transactionDetail.SenderWallet = transaction.Sender.Number
			transactionDetail.ReceiverName = transaction.Receiver.User.Name
			transactionDetail.ReceiverWallet = transaction.Receiver.Number
		}

		transactionDetails = append(transactionDetails, transactionDetail)
	}

	transactionResponse := res.TransactionPaginationResponses{
		Page:         query.Page,
		Total:        uint(len(transactionDetails)),
		Transactions: transactionDetails,
	}

	resp := res.Response{
		Message: "Get all transactions",
		Data:    transactionResponse,
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllHistoryUserTransactions implements interfaces.TransactionHandler.
func (handler *TransactionHandler) GetAllHistoryUserTransactions(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
	)

	getUserId, _ := c.Get("x-user-id")
	userId, _ := getUserId.(uint)

	transactions, err := handler.uscase.GetUserHistoryTransactions(ctx, userId)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Get all user transactions history",
		Data:    transactions,
	}
	c.JSON(http.StatusOK, resp)
}

// TopupTransaction implements interfaces.TransactionHandler.
func (handler *TransactionHandler) TopupTransaction(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
		request   *req.TopupRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		errBadRequest := &apperror.ErrFieldValidation{Message: utils.Validate(&request, err)}
		c.Error(errBadRequest)
		return
	}

	getUserId, _ := c.Get("x-user-id")
	userId, _ := getUserId.(uint)
	request.UserID = userId

	log.Println(request)
	transaction, err := handler.uscase.CreateTopupTransaction(ctx, request)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Success to top up balance",
		Data:    transaction,
	}

	c.JSON(http.StatusCreated, resp)
}

// TransferTransaction implements interfaces.TransactionHandler.
func (handler *TransactionHandler) TransferTransaction(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
		request   *req.TransferRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		errBadRequest := &apperror.ErrFieldValidation{Message: utils.Validate(&request, err)}
		c.Error(errBadRequest)
		return
	}

	getUserId, _ := c.Get("x-user-id")
	userId, _ := getUserId.(uint)
	request.SenderUserId = userId

	transaction, err := handler.uscase.CreateTransferTransaction(ctx, request)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Success to top up balance",
		Data:    transaction,
	}

	c.JSON(http.StatusCreated, resp)
}

func NewTransactionHandler(usecase interfaces.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{uscase: usecase}
}

var _ interfaces.TransactionHandler = &TransactionHandler{}
