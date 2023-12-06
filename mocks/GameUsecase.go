// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "assignment_4/entities/models"

	req "assignment_4/entities/payload/req"
)

// GameUsecase is an autogenerated mock type for the GameUsecase type
type GameUsecase struct {
	mock.Mock
}

// ChooseReward provides a mock function with given fields: ctx, chooseReward
func (_m *GameUsecase) ChooseReward(ctx context.Context, chooseReward *req.ChooseReward) (*models.Game, error) {
	ret := _m.Called(ctx, chooseReward)

	var r0 *models.Game
	if rf, ok := ret.Get(0).(func(context.Context, *req.ChooseReward) *models.Game); ok {
		r0 = rf(ctx, chooseReward)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Game)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *req.ChooseReward) error); ok {
		r1 = rf(ctx, chooseReward)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserCurrentChance provides a mock function with given fields: ctx, userId
func (_m *GameUsecase) GetUserCurrentChance(ctx context.Context, userId uint) (*models.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint) *models.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlayGame provides a mock function with given fields: ctx, game
func (_m *GameUsecase) PlayGame(ctx context.Context, game *models.Game) (*models.Game, error) {
	ret := _m.Called(ctx, game)

	var r0 *models.Game
	if rf, ok := ret.Get(0).(func(context.Context, *models.Game) *models.Game); ok {
		r0 = rf(ctx, game)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Game)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Game) error); ok {
		r1 = rf(ctx, game)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGameUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewGameUsecase creates a new instance of GameUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGameUsecase(t mockConstructorTestingTNewGameUsecase) *GameUsecase {
	mock := &GameUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
