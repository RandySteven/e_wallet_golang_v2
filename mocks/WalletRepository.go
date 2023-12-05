// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	interfaces "assignment_4/interfaces"
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "assignment_4/entities/models"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// BeginTrx provides a mock function with given fields: ctx
func (_m *WalletRepository) BeginTrx(ctx context.Context) interfaces.WalletRepository {
	ret := _m.Called(ctx)

	var r0 interfaces.WalletRepository
	if rf, ok := ret.Get(0).(func(context.Context) interfaces.WalletRepository); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.WalletRepository)
		}
	}

	return r0
}

// CommitOrRollback provides a mock function with given fields: ctx
func (_m *WalletRepository) CommitOrRollback(ctx context.Context) {
	_m.Called(ctx)
}

// FindAll provides a mock function with given fields: ctx
func (_m *WalletRepository) FindAll(ctx context.Context) ([]models.Wallet, error) {
	ret := _m.Called(ctx)

	var r0 []models.Wallet
	if rf, ok := ret.Get(0).(func(context.Context) []models.Wallet); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Wallet)
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

// GetById provides a mock function with given fields: ctx, id
func (_m *WalletRepository) GetById(ctx context.Context, id uint) (*models.Wallet, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uint) *models.Wallet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
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

// GetByNumber provides a mock function with given fields: ctx, number
func (_m *WalletRepository) GetByNumber(ctx context.Context, number string) (*models.Wallet, error) {
	ret := _m.Called(ctx, number)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Wallet); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserId provides a mock function with given fields: ctx, id
func (_m *WalletRepository) GetByUserId(ctx context.Context, id uint) (*models.Wallet, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uint) *models.Wallet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
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

// Save provides a mock function with given fields: ctx, entity
func (_m *WalletRepository) Save(ctx context.Context, entity *models.Wallet) (*models.Wallet, error) {
	ret := _m.Called(ctx, entity)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, *models.Wallet) *models.Wallet); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Wallet) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, entity
func (_m *WalletRepository) Update(ctx context.Context, entity *models.Wallet) (*models.Wallet, error) {
	ret := _m.Called(ctx, entity)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, *models.Wallet) *models.Wallet); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Wallet) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWalletRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepository(t mockConstructorTestingTNewWalletRepository) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
