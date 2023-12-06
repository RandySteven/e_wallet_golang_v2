package handlers

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

type GameHandler struct {
	usecase interfaces.GameUsecase
}

// ChooseBox implements interfaces.GameHandler.
func (handler *GameHandler) ChooseBox(c *gin.Context) {
	var (
		requestId    = uuid.NewString()
		ctx          = context.WithValue(c.Request.Context(), "request_id", requestId)
		chooseReward *req.ChooseReward
	)

	if err := c.ShouldBind(&chooseReward); err != nil {
		errBadRequest := &apperror.ErrFieldValidation{Message: utils.Validate(&chooseReward, err)}
		c.Error(errBadRequest)
		return
	}

	getGameId := c.Param("id")
	gameId, err := strconv.Atoi(getGameId)
	if err != nil {
		c.Error(&apperror.ErrInvalidFormat{Message: "invalid id format"})
		return
	}

	chooseReward.GameID = uint(gameId)

	game, err := handler.usecase.ChooseReward(ctx, chooseReward)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, game)
}

// PlayGame implements interfaces.GameHandler.
func (handler *GameHandler) PlayGame(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
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

	c.JSON(http.StatusCreated, game)
}

func (handler *GameHandler) CurrentUserChance(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)

	getUserId, _ := c.Get("x-user-id")
	userId, _ := getUserId.(uint)

	userChance, err := handler.usecase.GetUserCurrentChance(ctx, userId)
	if err != nil {
		c.Error(err)
		return
	}

	resp := res.Response{
		Message: "Get user current chance",
		Data:    userChance.Chance,
	}

	c.JSON(http.StatusOK, resp)
}

func NewGameHandler(usecase interfaces.GameUsecase) *GameHandler {
	return &GameHandler{usecase: usecase}
}

var _ interfaces.GameHandler = &GameHandler{}
