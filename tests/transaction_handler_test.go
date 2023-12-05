package tests

import (
	"assignment_4/handlers"
	middleware "assignment_4/middlewares"
	"assignment_4/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	transactionUsecase *mocks.TransactionUsecase
	transactionHandler *handlers.TransactionHandler
	router             *gin.Engine
}

func (suite *TransactionHandlerTestSuite) SetupSubTest() {
	suite.transactionUsecase = mocks.NewTransactionUsecase(suite.T())
	suite.transactionHandler = handlers.NewTransactionHandler(suite.transactionUsecase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestTransactionHandler(t *testing.T) {
	suite.Run(t, new(TransactionHandlerTestSuite))
}

func (suite *TransactionHandlerTestSuite) TestDoTransfer() {
	suite.Run("should return 201 after success transfer", func() {
		request := `{
			"to": "1000000000001",
			"amount": "20000.00",
			"description": "Ini ya duitnya"
		}`

		req, _ := http.NewRequest(http.MethodPost, "/v1/transfers", strings.NewReader(request))
		w := httptest.NewRecorder()

		suite.transactionUsecase.On(
			"CreateTransferTransaction", mock.Anything,
			mock.AnythingOfType("*req.TopupRequest"),
		).Return(mock.AnythingOfType("*models.Transaction"), nil)

		suite.router.POST("/v1/transfers", suite.transactionHandler.TransferTransaction)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusCreated, w.Code)
	})

	suite.Run("should return 400 failed to create transfer", func() {
		request := `{
			"to": "1000000000001",
			"amount": "20000.00",
			"description": "Ini ya duitnya"
		}`

		req, _ := http.NewRequest(http.MethodPost, "/v1/transfers", strings.NewReader(request))
		w := httptest.NewRecorder()

		suite.router.POST("/v1/transfers", suite.transactionHandler.TransferTransaction)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("should return 500 failed to create transfer", func() {

	})
}
