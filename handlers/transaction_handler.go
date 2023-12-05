package handlers

import (
	"assignment_4/entities"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"assignment_4/interfaces"
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

	sortBy := c.Query("sortBy")
	sort := c.DefaultQuery("sort", "asc")
	limit := c.DefaultQuery("limit", "25")
	page := c.DefaultQuery("page", "1")

	queryCondition := &entities.QueryCondition{
		SortedBy: sortBy,
		Sort:     sort,
		Limit:    limit,
		Page:     page,
	}

	transactions, err := handler.uscase.GetAllTransactionsRecords(ctx, queryCondition)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Get all transactions",
		Data:    transactions,
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
		c.Error(err)
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

	c.JSON(http.StatusOK, resp)
}

// TransferTransaction implements interfaces.TransactionHandler.
func (handler *TransactionHandler) TransferTransaction(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
		request   *req.TransferRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
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

	c.JSON(http.StatusOK, resp)
}

func NewTransactionHandler(usecase interfaces.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{uscase: usecase}
}

var _ interfaces.TransactionHandler = &TransactionHandler{}
