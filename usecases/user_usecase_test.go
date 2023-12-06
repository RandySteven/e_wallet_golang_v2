package usecases_test

import (
	"assignment_4/apperror"
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"assignment_4/enums"
	"assignment_4/mocks"
	"assignment_4/usecases"
	"assignment_4/utils"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	t.Run("success to register user", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		user := &models.User{
			Name:     "Randy Steven",
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
			Return(user, nil)

		ctx := context.Background()
		resultUser, err := userusecase.RegisterUser(ctx, user)

		assert.Nil(t, err)
		assert.Equal(t, resultUser.Name, user.Name)
	})

	t.Run("failed to register user", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		user := &models.User{
			Name:     "Randy Steven",
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		resultUser, err := userusecase.RegisterUser(ctx, user)

		assert.Error(t, err)
		assert.Nil(t, resultUser)
	})
}

func TestLoginUse(t *testing.T) {
	t.Run("success to login user", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		userLogin := &req.UserLoginRequest{
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		user := &models.User{
			Name:     "Randy Steven",
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		pass, _ := utils.HashPassword(user.Password)
		user.Password = pass
		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("GetByEmail", mock.Anything, userLogin.Email).
			Return(user, nil)

		ctx := context.Background()
		result, _ := userusecase.LoginUser(ctx, userLogin)

		assert.NotNil(t, result)
	})

	t.Run("failed to login user because password is not same", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		userLogin := &req.UserLoginRequest{
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		user := &models.User{
			Name:     "Randy Steven",
			Email:    "randy.steven@shopee.com",
			Password: "test_2345",
		}
		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("GetByEmail", mock.Anything, userLogin.Email).
			Return(user, nil)

		ctx := context.Background()
		_, err := userusecase.LoginUser(ctx, userLogin)

		assert.Error(t, err)
	})

	t.Run("failed to login user because error get query", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		userLogin := &req.UserLoginRequest{
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		user := &models.User{
			Name:     "Randy Steven",
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		pass, _ := utils.HashPassword(user.Password)
		user.Password = pass
		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("GetByEmail", mock.Anything, userLogin.Email).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := userusecase.LoginUser(ctx, userLogin)

		assert.Error(t, err)
	})
}

func TestGetUserDetail(t *testing.T) {
	t.Run("should return user detail success get user", func(t *testing.T) {
		userDetail := &res.UserDetail{
			Name:         "Randy Steven",
			Email:        "randy.steven@shopee.com",
			WalletNumber: "1000100101010",
			Balance:      decimal.NewFromInt(100000),
			Chance:       uint(0),
		}

		wallet := &models.Wallet{
			ID:      1,
			Number:  "1000100101010",
			Balance: decimal.NewFromInt(100000),
			User: &models.User{
				Name:  "Randy Steven",
				Email: "randy.steven@shopee.com",
			},
		}

		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		walletRepo.Mock.On("GetByUserId", mock.Anything, uint(1)).
			Return(wallet, nil)

		ctx := context.Background()
		res, _ := userusecase.GetUserDetail(ctx, 1)

		assert.Equal(t, userDetail, res)
	})

	t.Run("should return error user not found", func(t *testing.T) {

		wallet := &models.Wallet{}

		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		walletRepo.Mock.On("GetByUserId", mock.Anything, uint(1)).
			Return(wallet, nil)

		ctx := context.Background()
		_, err := userusecase.GetUserDetail(ctx, 1)

		assert.Errorf(t, err, "user not found")
	})

	t.Run("should return error from db", func(t *testing.T) {

		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		walletRepo.Mock.On("GetByUserId", mock.Anything, uint(1)).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		userDetail, err := userusecase.GetUserDetail(ctx, 1)

		assert.Error(t, err)
		assert.Nil(t, userDetail)
	})
}

func TestForgotPassword(t *testing.T) {
	t.Run("should return forgot password token", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		user := &models.User{
			ID:    1,
			Name:  "Randy Steven",
			Email: "randy.steven@shopee.com",
		}
		forgotPass := &req.ForgotPasswordRequest{
			Email: "randy.steven@shopee.com",
		}
		forgotPassToken := &models.ForgotPasswordToken{
			ResetToken:  uuid.NewString(),
			TokenExpiry: time.Now().Add(15 * time.Minute),
			IsValid:     true,
			UserID:      1,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("GetByEmail", mock.Anything, forgotPass.Email).
			Return(user, nil)

		forgotPassRepo.On("Save", mock.Anything, mock.AnythingOfType("*models.ForgotPasswordToken")).
			Return(forgotPassToken, nil)

		ctx := context.Background()
		expectForgotPassToken, _ := userusecase.ForgotPassword(ctx, forgotPass)

		assert.NotNil(t, expectForgotPassToken)
		assert.Equal(t, forgotPassToken.ResetToken, expectForgotPassToken.ResetToken)
	})

	t.Run("should return error invalid field email", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		forgotPass := &req.ForgotPasswordRequest{
			Email: "randy.steven@shopee.com",
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("GetByEmail", mock.Anything, forgotPass.Email).
			Return(nil, &apperror.ErrInvalidRequest{Field: enums.Email})

		ctx := context.Background()
		_, err := userusecase.ForgotPassword(ctx, forgotPass)

		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Errorf(t, err, "email invalid request")
	})

	t.Run("should return error while try to generate pass token", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		user := &models.User{
			ID:    1,
			Name:  "Randy Steven",
			Email: "randy.steven@shopee.com",
		}
		forgotPass := &req.ForgotPasswordRequest{
			Email: "randy.steven@shopee.com",
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)

		userRepo.On("GetByEmail", mock.Anything, forgotPass.Email).
			Return(user, nil)

		forgotPassRepo.On("Save", mock.Anything, mock.AnythingOfType("*models.ForgotPasswordToken")).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := userusecase.ForgotPassword(ctx, forgotPass)

		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Errorf(t, err, "mock error")
	})
}

func TestResetPassword(t *testing.T) {
	t.Run("should return success updated user", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		token := uuid.NewString()
		user := &models.User{
			ID:       1,
			Name:     "Randy Steven",
			Email:    "randy.steven@shopee.com",
			Password: "test_1234",
		}
		reset := &req.PasswordResetRequest{
			Email:       "randy.steven@shopee.com",
			NewPassword: "test_2345",
			Token:       token,
		}

		forgotPassToken := &models.ForgotPasswordToken{
			ID:          1,
			ResetToken:  token,
			TokenExpiry: time.Now().Add(time.Minute * 15),
			IsValid:     true,
			UserID:      1,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)
		forgotPassRepo.On("GetPasswordTokenByToken", mock.Anything, reset.Token).
			Return(forgotPassToken, nil)

		forgotPassRepo.On("UpdateUserPassword",
			mock.Anything,
			mock.AnythingOfType("*models.ForgotPasswordToken"),
			mock.Anything).
			Return(user, nil)

		ctx := context.Background()
		res, _ := userusecase.ResetPassword(ctx, reset)

		assert.NotEqual(t, "", res.Password)
	})

	t.Run("should return token already used", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		token := uuid.NewString()
		reset := &req.PasswordResetRequest{
			Email:       "randy.steven@shopee.com",
			NewPassword: "test_2345",
			Token:       token,
		}

		forgotPassToken := &models.ForgotPasswordToken{
			ID:          1,
			ResetToken:  token,
			TokenExpiry: time.Now().Add(time.Minute * 15),
			IsValid:     false,
			UserID:      1,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)
		forgotPassRepo.On("GetPasswordTokenByToken", mock.Anything, reset.Token).
			Return(forgotPassToken, nil)

		ctx := context.Background()
		_, err := userusecase.ResetPassword(ctx, reset)

		assert.Equal(t, "token is already used", err.Error())
	})

	t.Run("should return token already expired", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		token := uuid.NewString()
		reset := &req.PasswordResetRequest{
			Email:       "randy.steven@shopee.com",
			NewPassword: "test_2345",
			Token:       token,
		}

		forgotPassToken := &models.ForgotPasswordToken{
			ID:          1,
			ResetToken:  token,
			TokenExpiry: time.Now(),
			IsValid:     true,
			UserID:      1,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)
		forgotPassRepo.On("GetPasswordTokenByToken", mock.Anything, reset.Token).
			Return(forgotPassToken, nil)

		ctx := context.Background()
		_, err := userusecase.ResetPassword(ctx, reset)

		assert.Equal(t, "token already expired", err.Error())
	})

	t.Run("should return token invalid", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		token := uuid.NewString()
		reset := &req.PasswordResetRequest{
			Email:       "randy.steven@shopee.com",
			NewPassword: "test_2345",
			Token:       token,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)
		forgotPassRepo.On("GetPasswordTokenByToken", mock.Anything, reset.Token).
			Return(nil, nil)

		ctx := context.Background()
		_, err := userusecase.ResetPassword(ctx, reset)

		assert.Equal(t, "token is invalid", err.Error())
	})

	t.Run("should return token invalid", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		token := uuid.NewString()
		reset := &req.PasswordResetRequest{
			Email:       "randy.steven@shopee.com",
			NewPassword: "test_2345",
			Token:       token,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)
		forgotPassRepo.On("GetPasswordTokenByToken", mock.Anything, reset.Token).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := userusecase.ResetPassword(ctx, reset)

		assert.Equal(t, "mock error", err.Error())
	})

	t.Run("should return failed while updated user pass", func(t *testing.T) {
		walletRepo := &mocks.WalletRepository{}
		userRepo := &mocks.UserRepository{}
		forgotPassRepo := &mocks.ForgotPasswordRepository{}
		token := uuid.NewString()
		reset := &req.PasswordResetRequest{
			Email:       "randy.steven@shopee.com",
			NewPassword: "test_2345",
			Token:       token,
		}

		forgotPassToken := &models.ForgotPasswordToken{
			ID:          1,
			ResetToken:  token,
			TokenExpiry: time.Now().Add(time.Minute * 15),
			IsValid:     true,
			UserID:      1,
		}

		userusecase := usecases.NewUserUsecase(userRepo, walletRepo, forgotPassRepo)
		forgotPassRepo.On("GetPasswordTokenByToken", mock.Anything, reset.Token).
			Return(forgotPassToken, nil)

		forgotPassRepo.On("UpdateUserPassword",
			mock.Anything,
			mock.AnythingOfType("*models.ForgotPasswordToken"),
			mock.Anything).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := userusecase.ResetPassword(ctx, reset)

		assert.Equal(t, "mock error", err.Error())
	})

}
