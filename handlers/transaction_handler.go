package handlers

import (
	"assignment_4/entities/payload/req"
	"assignment_4/interfaces"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	uscase interfaces.TransactionUsecase
}

// TopupTransaction implements interfaces.TransactionHandler.
func (handler *TransactionHandler) TopupTransaction(c *gin.Context) {
	panic("unimplemented")
	var (
		// requestId = uuid.NewString()
		// ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
		request *req.TopupRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		return
	}

}

// TransferTransaction implements interfaces.TransactionHandler.
func (*TransactionHandler) TransferTransaction(c *gin.Context) {
	panic("unimplemented")
}

func NewTransactionHandler(usecase interfaces.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{uscase: usecase}
}

var _ interfaces.TransactionHandler = &TransactionHandler{}
