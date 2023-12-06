package handlers_test

import (
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	userUsecase *mocks.UserUsecase
	userHandler *handlers.UserHandler
	router      *gin.Engine
}

func (suite *UserHandlerTestSuite) SetupTest() {
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

func (suite *UserHandlerTestSuite) TestGetUserById() {
	userDetail := &res.UserDetail{}

	suite.userUsecase.On("GetUserDetail", mock.Anything, uint(1)).Return(userDetail, nil)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	suite.router.GET("/users/:id", suite.userHandler.GetUserById)
	suite.router.ServeHTTP(w, req)
	suite.T().Log("response : ", w.Body)

	suite.Assert().Equal(http.StatusOK, w.Code)
}

// Test LoginUser method
func (suite *UserHandlerTestSuite) TestLoginUser() {
	loginRequest := &req.UserLoginRequest{
		Email:    "randy.steven@gmail.com",
		Password: "test_1234",
	}

	userResponse := &res.UserLoginResponse{
		ID:    1,
		Name:  "Randy Steven",
		Email: "randy.steven@gmail.com",
		Token: "lalalala",
	}

	suite.userUsecase.On("LoginUser", mock.Anything, mock.AnythingOfType("*req.UserLoginRequest")).Return(userResponse, nil)

	loginRequestBody, err := json.Marshal(loginRequest)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()

	suite.router.POST("/login", suite.userHandler.LoginUser)
	suite.router.ServeHTTP(w, req)

	suite.T().Log(w.Body)
	suite.Equal(http.StatusOK, w.Code)

}

func (suite *UserHandlerTestSuite) TestLoginFailed() {
	loginRequest := &req.UserLoginRequest{
		Password: "test_1234",
	}

	loginRequestBody, err := json.Marshal(loginRequest)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()

	suite.router.POST("/login", suite.userHandler.LoginUser)
	suite.router.ServeHTTP(w, req)

	suite.T().Log(w.Body)
	suite.Equal(http.StatusBadRequest, w.Code)
}

// Test RegisterUser method
func (suite *UserHandlerTestSuite) TestRegisterUser() {
	registerRequest := &req.UserRegisterRequest{
		Name:     "Randy Steven",
		Email:    "randy.steven@gmail.com",
		Password: "test_1234",
	}

	suite.userUsecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).Return(&users[0], nil)

	registerRequestBody, err := json.Marshal(registerRequest)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()

	suite.router.POST("/register", suite.userHandler.RegisterUser)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	suite.T().Log(w.Body)
}

func (suite *UserHandlerTestSuite) TestFailedRegisterUser() {
	registerRequest := &req.UserRegisterRequest{
		Name:  "Randy Steven",
		Email: "randy.steven@gmail.com",
	}

	registerRequestBody, err := json.Marshal(registerRequest)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()

	suite.router.POST("/register", suite.userHandler.RegisterUser)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	suite.T().Log(w.Body)
}

func (suite *UserHandlerTestSuite) TestFailedRegisterUserInternalServerError() {
	registerRequest := &req.UserRegisterRequest{
		Name:     "Randy Steven",
		Email:    "randy.steven@gmail.com",
		Password: "test_1234",
	}

	suite.userUsecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
		Return(nil, errors.New("mock error"))

	registerRequestBody, _ := json.Marshal(registerRequest)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(registerRequestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/register", suite.userHandler.RegisterUser)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)

}
