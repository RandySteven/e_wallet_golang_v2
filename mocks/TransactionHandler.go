// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// TransactionHandler is an autogenerated mock type for the TransactionHandler type
type TransactionHandler struct {
	mock.Mock
}

// GetAllHistoryUserTransactions provides a mock function with given fields: c
func (_m *TransactionHandler) GetAllHistoryUserTransactions(c *gin.Context) {
	_m.Called(c)
}

// GetAllTransactionsRecords provides a mock function with given fields: c
func (_m *TransactionHandler) GetAllTransactionsRecords(c *gin.Context) {
	_m.Called(c)
}

// TopupTransaction provides a mock function with given fields: c
func (_m *TransactionHandler) TopupTransaction(c *gin.Context) {
	_m.Called(c)
}

// TransferTransaction provides a mock function with given fields: c
func (_m *TransactionHandler) TransferTransaction(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewTransactionHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionHandler creates a new instance of TransactionHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionHandler(t mockConstructorTestingTNewTransactionHandler) *TransactionHandler {
	mock := &TransactionHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
