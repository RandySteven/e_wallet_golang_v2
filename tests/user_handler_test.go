package tests

import (
	"assignment_4/entities/models"
	"assignment_4/handlers"
	middleware "assignment_4/middlewares"
	"assignment_4/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	userUsecase *mocks.UserUsecase
	userHandler *handlers.UserHandler
	router      *gin.Engine
}

func (suite *UserHandlerTestSuite) SetupSubTest() {
	suite.userUsecase = mocks.NewUserUsecase(suite.T())
	suite.userHandler = handlers.NewUserHandler(suite.userUsecase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

var users = []models.User{
	{
		ID:    1,
		Name:  "Randy Steven",
		Email: "randy.steven@gmail.com",
	},
	{
		ID:    2,
		Name:  "Matthew Alfredo",
		Email: "matthew.alfredo@gmail.com",
	},
}

func (suite *UserHandlerTestSuite) TestRegisterUser() {
	suite.Run("should return 201 success to create user and wallet", func() {
		requestBody := `{
			"name": "Matthew Alfredo",
			"email": "matthew.alfredo@gmail.com",
			"password": "test_1234"
		  }`
		req, _ := http.NewRequest("POST", "/v1/register", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		suite.userUsecase.
			On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
			Return(&users[0], nil)

		suite.router.POST("/v1/register", suite.userHandler.RegisterUser)
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusCreated, w.Code)
	})
}
