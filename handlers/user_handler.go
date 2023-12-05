package handlers

import (
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"assignment_4/interfaces"
	"assignment_4/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

// LoginUser implements interfaces.UserHandler.
func (handler *UserHandler) LoginUser(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		login     *req.UserLoginRequest
	)
	if err := c.ShouldBind(&login); err != nil {
		return
	}
	userRes, err := handler.usecase.LoginUser(ctx, login)
	if err != nil {
		return
	}
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pass, err := utils.HashPassword(register.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := &models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: pass,
	}
	user, err = handler.usecase.RegisterUser(ctx, user)
	if err != nil {
		return
	}

	resp := res.Response{
		Message: "Success created user",
		Data:    user,
	}
	c.JSON(http.StatusCreated, resp)
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ interfaces.UserHandler = &UserHandler{}
