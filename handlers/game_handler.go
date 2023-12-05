package handlers

import (
	"assignment_4/interfaces"

	"github.com/gin-gonic/gin"
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
	// var (
	// requestId = uuid.NewString()
	// ctx       = context.WithValue(c.Request.Context(), "requestId", requestId)
	// )

	// getUserId, _ := c.Get("x-user-id")
	// userId, _ := getUserId.(uint)

}

func NewGameHandler(usecase interfaces.GameUsecase) *GameHandler {
	return &GameHandler{usecase: usecase}
}

var _ interfaces.GameHandler = &GameHandler{}
