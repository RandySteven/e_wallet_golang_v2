package handlers_rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/res"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	uscase interfaces.TransactionUsecase
}

// GetAllTransactionsRecords implements interfaces.TransactionHandler.
func (handler *TransactionHandler) GetAllTransactionsRecords(c *gin.Context) {
	var (
		requestId    = uuid.NewString()
		ctx          = context.WithValue(c.Request.Context(), "requestId", requestId)
		getUserId, _ = c.Get("x-user-id")
		userId, _    = getUserId.(uint)
	)

	var query entities.QueryCondition
	c.ShouldBindQuery(&query)

	if query.StartDate != "" || query.EndDate != "" {
		startDate, err := time.Parse("2006-01-02", query.StartDate)
		if err != nil {
			c.Error(&apperror.ErrInvalidFormat{Message: "Date must in format YYYY-MM-dd"})
			return
		}
		endDate, err := time.Parse("2006-01-02", query.EndDate)
		if err != nil {
			c.Error(&apperror.ErrInvalidFormat{Message: "Date must in format YYYY-MM-dd"})
			return
		}
		query.StartDate = startDate.Format("2006-01-02")
		query.EndDate = endDate.Format("2006-01-02")
	}

	transactionResponse, err := handler.uscase.GetAllTransactionsRecords(ctx, &query, userId)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Get all transactions",
		Data:    transactionResponse,
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

	transactionDetail := res.TransactionResponse(transaction)
	resp := res.Response{
		Message: "Success to top up balance",
		Data:    transactionDetail,
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

	transactionRes := res.TransactionResponse(transaction)
	resp := res.Response{
		Message: "Success to top up balance",
		Data:    transactionRes,
	}

	c.JSON(http.StatusCreated, resp)
}

func NewTransactionHandler(usecase interfaces.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{uscase: usecase}
}

var _ interfaces.TransactionHandler = &TransactionHandler{}
