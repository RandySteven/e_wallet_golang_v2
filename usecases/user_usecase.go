package usecases

import (
	"assignment_4/apperror"
	"assignment_4/auth"
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"assignment_4/enums"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo       interfaces.UserRepository
	walletRepo     interfaces.WalletRepository
	forgotPassRepo interfaces.ForgotPasswordRepository
}

// ResetPassword implements interfaces.UserUsecase.
func (usecase *userUsecase) ResetPassword(ctx context.Context, reset *req.PasswordResetRequest) (*models.User, error) {
	token, err := usecase.forgotPassRepo.GetPasswordTokenByToken(ctx, reset.Token)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, &apperror.ErrTokenInvalid{}
	}

	if !token.IsValid {
		return nil, &apperror.ErrTokenAlreadyUsed{}
	}

	if time.Now().After(token.TokenExpiry) {
		return nil, &apperror.ErrTokenExpired{}
	}

	pass, err := utils.HashPassword(reset.NewPassword)
	if err != nil {
		return nil, err
	}

	user, err := usecase.forgotPassRepo.UpdateUserPassword(ctx, token, pass)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ForgotPassword implements interfaces.UserUsecase.
func (usecase *userUsecase) ForgotPassword(ctx context.Context, forgot *req.ForgotPasswordRequest) (*models.ForgotPasswordToken, error) {
	user, err := usecase.userRepo.GetByEmail(ctx, forgot.Email)
	if err != nil || user == nil {
		return nil, &apperror.ErrInvalidRequest{Field: enums.Email}
	}

	currentTime := time.Now()
	forgotPassword := &models.ForgotPasswordToken{
		ResetToken:  uuid.NewString(),
		TokenExpiry: currentTime.Add(15 * time.Minute),
		IsValid:     true,
		UserID:      user.ID,
	}

	forgotPassword, err = usecase.forgotPassRepo.Save(ctx, forgotPassword)
	if err != nil {
		return nil, err
	}

	return forgotPassword, nil
}

// GetUserDetail implements interfaces.UserUsecase.
func (usecase *userUsecase) GetUserDetail(ctx context.Context, id uint) (*res.UserDetail, error) {
	wallet, err := usecase.walletRepo.GetByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	userDetail := &res.UserDetail{
		Name:         wallet.User.Name,
		Email:        wallet.User.Email,
		WalletNumber: wallet.Number,
		Balance:      wallet.Balance,
		Chance:       wallet.User.Chance,
	}

	return userDetail, nil
}

// LoginUser implements interfaces.UserUsecase.
func (usecase *userUsecase) LoginUser(ctx context.Context, login *req.UserLoginRequest) (*res.UserLoginResponse, error) {
	user, err := usecase.userRepo.GetByEmail(ctx, login.Email)
	if err != nil {
		return nil, err
	}
	isPassValid := utils.IsPasswordValid(user.Password, login.Password)
	if !isPassValid {
		return nil, err
	}
	expTime := time.Now().Add(time.Hour * 1)
	claims := &auth.JWTClaim{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APPLICATION",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(auth.JWT_KEY)
	log.Println(string(auth.JWT_KEY))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	userResp := res.UserLoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	return &userResp, nil
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := usecase.userRepo.RegisterUser(ctx, user)
	return user, err
}

func NewUserUsecase(userRepo interfaces.UserRepository,
	walletRepo interfaces.WalletRepository,
	forgotPassRepo interfaces.ForgotPasswordRepository) *userUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		walletRepo:     walletRepo,
		forgotPassRepo: forgotPassRepo,
	}
}

var _ interfaces.UserUsecase = &userUsecase{}
