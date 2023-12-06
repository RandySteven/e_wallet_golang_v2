package handlers_test

import (
	"assignment_4/apperror"
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/handlers"
	middleware "assignment_4/middlewares"
	"assignment_4/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GameHandlerTestSuite struct {
	suite.Suite
	gameUsecase *mocks.GameUsecase
	gameHandler *handlers.GameHandler
	router      *gin.Engine
}

func (suite *GameHandlerTestSuite) SetupTest() {
	suite.gameUsecase = mocks.NewGameUsecase(suite.T())
	suite.gameHandler = handlers.NewGameHandler(suite.gameUsecase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestGameHandler(t *testing.T) {
	suite.Run(t, new(GameHandlerTestSuite))
}

var games = []models.Game{
	{
		ID:       1,
		UserID:   1,
		BoxID1:   2,
		BoxID2:   3,
		BoxID3:   1,
		BoxID4:   12,
		BoxID5:   9,
		BoxID6:   4,
		BoxID7:   5,
		BoxID8:   8,
		BoxID9:   11,
		WinBoxID: 11,
	},
	{
		ID:       2,
		UserID:   1,
		BoxID1:   2,
		BoxID2:   3,
		BoxID3:   1,
		BoxID4:   12,
		BoxID5:   9,
		BoxID6:   4,
		BoxID7:   5,
		BoxID8:   8,
		BoxID9:   11,
		WinBoxID: 11,
	},
}

func (suite *GameHandlerTestSuite) TestPlayGame() {
	req, _ := http.NewRequest(http.MethodPost, "/v1/games", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.gameUsecase.On("PlayGame", mock.Anything, mock.AnythingOfType("*models.Game")).
		Return(&games[0], nil)

	suite.router.POST("/v1/games", suite.gameHandler.PlayGame)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
}

func (suite *GameHandlerTestSuite) TestPlayGameInternalServer() {
	req, _ := http.NewRequest(http.MethodPost, "/v1/games", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.gameUsecase.On("PlayGame", mock.Anything, mock.AnythingOfType("*models.Game")).
		Return(nil, errors.New("mock error"))

	suite.router.POST("/v1/games", suite.gameHandler.PlayGame)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *GameHandlerTestSuite) TestChooseGameRewards() {
	chooseReward := &req.ChooseReward{
		BoxID: 12,
	}

	chooseRewardRequest, _ := json.Marshal(chooseReward)
	req, _ := http.NewRequest(http.MethodPut, "/v1/games/1", bytes.NewBuffer(chooseRewardRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.gameUsecase.On("ChooseReward", mock.Anything, mock.AnythingOfType("*req.ChooseReward")).
		Return(&games[0], nil)

	suite.router.PUT("/v1/games/:id", suite.gameHandler.ChooseBox)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *GameHandlerTestSuite) TestChooseGameRewardsBadRequest() {
	chooseReward := &req.ChooseReward{}

	chooseRewardRequest, _ := json.Marshal(chooseReward)
	req, _ := http.NewRequest(http.MethodPut, "/v1/games/1", bytes.NewBuffer(chooseRewardRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.PUT("/v1/games/:id", suite.gameHandler.ChooseBox)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *GameHandlerTestSuite) TestChooseGameRewardsBadRequestGameId() {
	chooseReward := &req.ChooseReward{
		BoxID: 1,
	}

	chooseRewardRequest, _ := json.Marshal(chooseReward)
	req, _ := http.NewRequest(http.MethodPut, "/v1/games/A", bytes.NewBuffer(chooseRewardRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.PUT("/v1/games/:id", suite.gameHandler.ChooseBox)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *GameHandlerTestSuite) TestChooseGameRewardsGameIdNotFound() {
	chooseReward := &req.ChooseReward{
		BoxID: 12,
	}

	chooseRewardRequest, _ := json.Marshal(chooseReward)
	req, _ := http.NewRequest(http.MethodPut, "/v1/games/1", bytes.NewBuffer(chooseRewardRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.gameUsecase.On("ChooseReward", mock.Anything, mock.AnythingOfType("*req.ChooseReward")).
		Return(nil, &apperror.ErrDataNotFound{Data: "game"})

	suite.router.PUT("/v1/games/:id", suite.gameHandler.ChooseBox)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *GameHandlerTestSuite) TestChooseGameRewardsInternalServerError() {
	chooseReward := &req.ChooseReward{
		BoxID: 12,
	}

	chooseRewardRequest, _ := json.Marshal(chooseReward)
	req, _ := http.NewRequest(http.MethodPut, "/v1/games/1", bytes.NewBuffer(chooseRewardRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.gameUsecase.On("ChooseReward", mock.Anything, mock.AnythingOfType("*req.ChooseReward")).
		Return(nil, errors.New("mock error"))

	suite.router.PUT("/v1/games/:id", suite.gameHandler.ChooseBox)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *GameHandlerTestSuite) TestCurrentUserChance() {
	var userId uint = 0
	req, _ := http.NewRequest(http.MethodGet, "/v1/games/chances", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	user := &models.User{
		ID:     1,
		Name:   "Randy Steven",
		Email:  "randy.steven@shopee.com",
		Chance: 1,
	}

	suite.gameUsecase.On("GetUserCurrentChance", mock.Anything, userId).
		Return(user, nil)

	suite.router.GET("/v1/games/chances", suite.gameHandler.CurrentUserChance)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *GameHandlerTestSuite) TestCurrentUserChanceError() {
	var userId uint = 0
	req, _ := http.NewRequest(http.MethodGet, "/v1/games/chances", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.gameUsecase.On("GetUserCurrentChance", mock.Anything, userId).
		Return(nil, errors.New("mock error"))

	suite.router.GET("/v1/games/chances", suite.gameHandler.CurrentUserChance)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}
