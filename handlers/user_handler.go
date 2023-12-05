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
func (*UserHandler) LoginUser(c *gin.Context) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserHandler.
func (handler *UserHandler) RegisterUser(c *gin.Context) {
	var register *req.UserRegisterRequest

	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)

	if err := c.ShouldBind(&register); err != nil {
		return
	}

	pass, err := utils.HashPassword(register.Password)
	if err != nil {
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
