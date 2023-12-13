package handlers_rest

import (
	"context"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/res"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

// ForgotPassword implements interfaces.UserHandler.
func (handler *UserHandler) ForgotPassword(c *gin.Context) {
	var (
		requestId      = uuid.NewString()
		ctx            = context.WithValue(c.Request.Context(), "request_id", requestId)
		forgotPassword *req.ForgotPasswordRequest
	)

	if err := c.ShouldBind(&forgotPassword); err != nil {
		errorBad := &apperror.ErrFieldValidation{Message: utils.Validate(&forgotPassword, err)}

		c.Error(errorBad)
		return
	}

	forgotPasswordToken, err := handler.usecase.ForgotPassword(ctx, forgotPassword)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Forgot password token",
		Data:    forgotPasswordToken,
	}

	c.JSON(http.StatusOK, resp)
}

// ResetPassword implements interfaces.UserHandler.
func (handler *UserHandler) ResetPassword(c *gin.Context) {
	var (
		requestId   = uuid.NewString()
		ctx         = context.WithValue(c.Request.Context(), "request_id", requestId)
		newPassword *req.PasswordResetRequest
	)

	if err := c.ShouldBind(&newPassword); err != nil {
		errBadRequest := &apperror.ErrFieldValidation{Message: utils.Validate(&newPassword, err)}
		c.Error(errBadRequest)
		return
	}

	newPassword.Token = c.Query("token")

	_, err := handler.usecase.ResetPassword(ctx, newPassword)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Success updated new password",
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserById implements interfaces.UserHandler.
func (handler *UserHandler) GetUserById(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.Error(&apperror.ErrInvalidFormat{Message: "id invalid format"})
		return
	}

	userDetail, err := handler.usecase.GetUserDetail(ctx, uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Success get user",
		Data:    userDetail,
	}

	c.JSON(http.StatusOK, resp)
}

// LoginUser implements interfaces.UserHandler.
func (handler *UserHandler) LoginUser(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		login     *req.UserLoginRequest
	)
	if err := c.ShouldBind(&login); err != nil {
		errBadRequest := &apperror.ErrFieldValidation{Message: utils.Validate(&login, err)}
		c.Error(errBadRequest)
		return
	}
	userRes, err := handler.usecase.LoginUser(ctx, login)
	if err != nil {
		c.Error(err)
		return
	}
	c.Header("Authorization", "Bearer "+userRes.Token)
	resp := &res.Response{
		Message: "Success to login user",
		Data:    userRes,
	}
	c.JSON(http.StatusOK, resp)
}

// RegisterUser implements interfaces.UserHandler.
func (handler *UserHandler) RegisterUser(c *gin.Context) {
	var register *req.UserRegisterRequest

	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)

	if err := c.ShouldBind(&register); err != nil {
		errorBad := &apperror.ErrFieldValidation{Message: utils.Validate(&register, err)}

		c.Error(errorBad)
		return
	}

	pass, _ := utils.HashPassword(register.Password)

	user := &models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: pass,
	}
	user, err := handler.usecase.RegisterUser(ctx, user)
	if err != nil {
		c.Error(err)
		return
	}

	userResp := res.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	resp := res.Response{
		Message: "Success created user",
		Data:    userResp,
	}
	c.JSON(http.StatusCreated, resp)
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ interfaces.UserHandler = &UserHandler{}
