package handlers

import (
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GameHandler struct {
	usecase interfaces.GameUsecase
}

// ChooseBox implements interfaces.GameHandler.
func (handler *GameHandler) ChooseBox(c *gin.Context) {
	panic("unimplemented")
}

// PlayGame implements interfaces.GameHandler.
func (handler *GameHandler) PlayGame(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
	)

	getUserId, _ := c.Get("x-user-id")
	userId, _ := getUserId.(uint)

	game := &models.Game{
		UserID: userId,
	}
	game, err := handler.usecase.PlayGame(ctx, game)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, game)
}

func NewGameHandler(usecase interfaces.GameUsecase) *GameHandler {
	return &GameHandler{usecase: usecase}
}

var _ interfaces.GameHandler = &GameHandler{}
