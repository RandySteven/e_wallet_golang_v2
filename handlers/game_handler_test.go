package handlers_test

import (
	"assignment_4/entities/models"
	"assignment_4/handlers"
	middleware "assignment_4/middlewares"
	"assignment_4/mocks"
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
