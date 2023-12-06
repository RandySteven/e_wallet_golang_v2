package tests

import (
	"assignment_4/entities/models"
	"assignment_4/handlers"
	middleware "assignment_4/middlewares"
	"assignment_4/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	transactionUsecase *mocks.TransactionUsecase
	transactionHandler *handlers.TransactionHandler
	router             *gin.Engine
}

func (suite *TransactionHandlerTestSuite) SetupTest() {
	suite.transactionUsecase = mocks.NewTransactionUsecase(suite.T())
	suite.transactionHandler = handlers.NewTransactionHandler(suite.transactionUsecase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestTransactionHandler(t *testing.T) {
	suite.Run(t, new(TransactionHandlerTestSuite))
}

var transactions = []models.Transaction{
	{
		ID:             1,
		SenderID:       1,
		ReceiverID:     2,
		Amount:         decimal.NewFromInt(50000),
		Description:    "",
		SourceOfFundID: 5,
	},
	{
		ID:             2,
		SenderID:       1,
		ReceiverID:     1,
		Amount:         decimal.NewFromInt(50000),
		Description:    "",
		SourceOfFundID: 1,
	},
}

func (suite *TransactionHandlerTestSuite) TestSuccessTransfer() {
	request := `{
			"to": "1000000000001",
			"amount": "20000.00",
			"description": "Ini ya duitnya"
		}`

	req, _ := http.NewRequest(http.MethodPost, "/v1/transfers", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.transactionUsecase.On(
		"CreateTransferTransaction", mock.Anything,
		mock.AnythingOfType("*req.TransferRequest"),
	).Return(&transactions[0], nil)

	suite.router.POST("/v1/transfers", suite.transactionHandler.TransferTransaction)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

}

// func (suite *TransactionHandlerTestSuite) TestFailedBadRequestTransfer() {
// 	request := `{
// 		"to": "1000000000001",
// 		"description": "Ini ya duitnya"
// 	}`

// 	req, _ := http.NewRequest(http.MethodPost, "/v1/transfers", strings.NewReader(request))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	suite.router.POST("/v1/transfers", suite.transactionHandler.TransferTransaction)
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusBadRequest, w.Code)
// }

func (suite *TransactionHandlerTestSuite) TestInternalServerErrorTransfer() {
	request := `{
		"to": "1000000000001",
		"amount": "20000.00",
		"description": "Ini ya duitnya"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/v1/transfers", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.transactionUsecase.On(
		"CreateTransferTransaction", mock.Anything,
		mock.AnythingOfType("*req.TransferRequest"),
	).Return(nil, errors.New("mock error"))

	suite.router.POST("/v1/transfers", suite.transactionHandler.TransferTransaction)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestSuccessDoTopup() {
	request := `{
			"amount": "10000",
			"source_of_fund": "Bank Transfer"
		}`

	req, _ := http.NewRequest(http.MethodPost, "/v1/topups", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.transactionUsecase.
		On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*req.TopupRequest")).
		Return(&transactions[1], nil)

	suite.router.POST("/v1/topups", suite.transactionHandler.TopupTransaction)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

}

func (suite *TransactionHandlerTestSuite) TestFailedTopup() {
	request := `{
			"amount": "10000"
		}`

	req, _ := http.NewRequest(http.MethodPost, "/v1/topups", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/v1/topups", suite.transactionHandler.TopupTransaction)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestDBErrorTopup() {
	request := `{
		"amount": "10000",
		"source_of_fund": "Bank Transfer"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/v1/topups", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.transactionUsecase.
		On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*req.TopupRequest")).
		Return(nil, errors.New("mock error"))

	suite.router.POST("/v1/topups", suite.transactionHandler.TopupTransaction)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}
