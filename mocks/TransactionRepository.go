// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entities "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities"
	context "context"

	interfaces "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"

	mock "github.com/stretchr/testify/mock"

	models "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// BeginTrx provides a mock function with given fields: ctx
func (_m *TransactionRepository) BeginTrx(ctx context.Context) interfaces.TransactionRepository {
	ret := _m.Called(ctx)

	var r0 interfaces.TransactionRepository
	if rf, ok := ret.Get(0).(func(context.Context) interfaces.TransactionRepository); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.TransactionRepository)
		}
	}

	return r0
}

// CommitOrRollback provides a mock function with given fields: ctx
func (_m *TransactionRepository) CommitOrRollback(ctx context.Context) {
	_m.Called(ctx)
}

// Count provides a mock function with given fields: ctx
func (_m *TransactionRepository) Count(ctx context.Context) (uint, error) {
	ret := _m.Called(ctx)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context) uint); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTopupTransaction provides a mock function with given fields: ctx, transaction
func (_m *TransactionRepository) CreateTopupTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	ret := _m.Called(ctx, transaction)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *models.Transaction) *models.Transaction); ok {
		r0 = rf(ctx, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Transaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransferTransaction provides a mock function with given fields: ctx, transaction
func (_m *TransactionRepository) CreateTransferTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	ret := _m.Called(ctx, transaction)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *models.Transaction) *models.Transaction); ok {
		r0 = rf(ctx, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Transaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: ctx
func (_m *TransactionRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	ret := _m.Called(ctx)

	var r0 []models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context) []models.Transaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTransactions provides a mock function with given fields: ctx, query, walletId
func (_m *TransactionRepository) GetAllTransactions(ctx context.Context, query *entities.QueryCondition, walletId uint) ([]models.Transaction, error) {
	ret := _m.Called(ctx, query, walletId)

	var r0 []models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *entities.QueryCondition, uint) []models.Transaction); ok {
		r0 = rf(ctx, query, walletId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.QueryCondition, uint) error); ok {
		r1 = rf(ctx, query, walletId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *TransactionRepository) GetById(ctx context.Context, id uint) (*models.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint) *models.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionCountBasedUserId provides a mock function with given fields: ctx, id
func (_m *TransactionRepository) GetTransactionCountBasedUserId(ctx context.Context, id uint) (uint, error) {
	ret := _m.Called(ctx, id)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, uint) uint); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByWalletId provides a mock function with given fields: ctx, walletId
func (_m *TransactionRepository) GetTransactionsByWalletId(ctx context.Context, walletId uint) ([]models.Transaction, error) {
	ret := _m.Called(ctx, walletId)

	var r0 []models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint) []models.Transaction); ok {
		r0 = rf(ctx, walletId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, walletId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, entity
func (_m *TransactionRepository) Save(ctx context.Context, entity *models.Transaction) (*models.Transaction, error) {
	ret := _m.Called(ctx, entity)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *models.Transaction) *models.Transaction); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Transaction) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, entity
func (_m *TransactionRepository) Update(ctx context.Context, entity *models.Transaction) (*models.Transaction, error) {
	ret := _m.Called(ctx, entity)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *models.Transaction) *models.Transaction); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Transaction) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
