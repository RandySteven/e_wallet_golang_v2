package handlers_rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/res"
	rest "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/handlers/rest"
	middleware "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/middlewares"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/mocks"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	transactionUsecase *mocks.TransactionUsecase
	transactionHandler *rest.TransactionHandler
	router             *gin.Engine
}

func (suite *TransactionHandlerTestSuite) SetupTest() {
	suite.transactionUsecase = mocks.NewTransactionUsecase(suite.T())
	suite.transactionHandler = rest.NewTransactionHandler(suite.transactionUsecase)
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

func (suite *TransactionHandlerTestSuite) TestFailedBadRequestTransfer() {
	request := `{
		"amount": "50000000",
		"description": "Ini ya duitnya"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/v1/transfers", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/v1/transfers", suite.transactionHandler.TransferTransaction)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

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

func (suite *TransactionHandlerTestSuite) TestGetAllListTransactions() {
	req, _ := http.NewRequest(http.MethodGet, "/v1/transactions", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	transactionPage := &res.TransactionPaginationResponses{
		Page:         "1",
		Total:        1,
		Transactions: []res.TransactionDetailResponse{},
	}

	suite.transactionUsecase.
		On("GetAllTransactionsRecords", mock.Anything, mock.AnythingOfType("*entities.QueryCondition"), uint(0)).
		Return(transactionPage, nil)

	suite.router.GET("/v1/transactions", suite.transactionHandler.GetAllTransactionsRecords)
	suite.router.ServeHTTP(w, req)
	suite.T().Log(w.Body)
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetAllListTransactionsFilterByDate() {
	req, _ := http.NewRequest(http.MethodGet, "/v1/transactions", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	q := req.URL.Query()
	q.Add("start_date", "2023-01-02")
	q.Add("end_date", "2023-01-03")
	req.URL.RawQuery = q.Encode()

	transactionPage := &res.TransactionPaginationResponses{
		Page:         "1",
		Total:        1,
		Transactions: []res.TransactionDetailResponse{},
	}

	suite.transactionUsecase.
		On("GetAllTransactionsRecords", mock.Anything, mock.AnythingOfType("*entities.QueryCondition"), uint(0)).
		Return(transactionPage, nil)

	suite.router.GET("/v1/transactions", suite.transactionHandler.GetAllTransactionsRecords)
	suite.router.ServeHTTP(w, req)
	suite.T().Log(w.Body)
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetAllListTransactionsFailedInvalidStartDate() {
	req, _ := http.NewRequest(http.MethodGet, "/v1/transactions", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	q := req.URL.Query()
	q.Add("start_date", "ABCDEF")
	req.URL.RawQuery = q.Encode()

	suite.router.GET("/v1/transactions", suite.transactionHandler.GetAllTransactionsRecords)
	suite.router.ServeHTTP(w, req)
	suite.T().Log(w.Body)
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetAllListTransactionsFailedInvalidEndDate() {
	req, _ := http.NewRequest(http.MethodGet, "/v1/transactions", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	q := req.URL.Query()
	q.Add("start_date", "2023-01-02")
	q.Add("end_date", "ABCDEF")
	req.URL.RawQuery = q.Encode()

	suite.router.GET("/v1/transactions", suite.transactionHandler.GetAllTransactionsRecords)
	suite.router.ServeHTTP(w, req)
	suite.T().Log(w.Body)
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetAllListTransactionsInternalServerError() {
	req, _ := http.NewRequest(http.MethodGet, "/v1/transactions", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.transactionUsecase.
		On("GetAllTransactionsRecords", mock.Anything, mock.AnythingOfType("*entities.QueryCondition"), uint(0)).
		Return(nil, errors.New("mock error"))

	suite.router.GET("/v1/transactions", suite.transactionHandler.GetAllTransactionsRecords)
	suite.router.ServeHTTP(w, req)
	suite.T().Log(w.Body)
	suite.Equal(http.StatusInternalServerError, w.Code)
}
