package tests

import (
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"assignment_4/handlers"
	middleware "assignment_4/middlewares"
	"assignment_4/mocks"
	"bytes"
	"encoding/json"
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

func (suite *UserHandlerTestSuite) TestGetUserById() {
	userDetail := &res.UserDetail{}

	suite.userUsecase.On("GetUserDetail", mock.Anything, uint(1)).Return(userDetail, nil)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	suite.router.GET("/users/:id", suite.userHandler.GetUserById)
	suite.router.ServeHTTP(w, req)

	suite.Assert().Equal(http.StatusOK, w.Code)

	var respBody map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &respBody)

	suite.Assert().Equal("Success get user", respBody["Message"])
	suite.Assert().Equal(userDetail, respBody["Data"])
}

// Test LoginUser method
func (suite *UserHandlerTestSuite) TestLoginUser() {
	// Mock data
	loginRequest := &req.UserLoginRequest{
		Email:    "test@example.com",
		Password: "password123",
		// Add other fields as needed
	}

	userResponse := &models.User{
		ID:   1,
		Name: "John Doe",
		// Add other fields as needed
	}

	// Set up expectations
	suite.userUsecase.On("LoginUser", mock.Anything, loginRequest).Return(userResponse, nil)

	// Convert loginRequest to JSON
	loginRequestBody, err := json.Marshal(loginRequest)
	assert.NoError(suite.T(), err)

	// Create a request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request to the router
	suite.router.POST("/login", suite.userHandler.LoginUser)
	suite.router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Unmarshal response body
	var respBody map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.NoError(suite.T(), err)

	// Assert response body
	assert.Equal(suite.T(), "Success to login user", respBody["Message"])
	assert.Equal(suite.T(), userResponse, respBody["Data"])
}

// Test RegisterUser method
func (suite *UserHandlerTestSuite) TestRegisterUser() {
	// Mock data
	registerRequest := &req.UserRegisterRequest{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Set up expectations
	suite.userUsecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).Return(&users[0], nil)

	// Convert registerRequest to JSON
	registerRequestBody, err := json.Marshal(registerRequest)
	assert.NoError(suite.T(), err)

	// Create a request
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request to the router
	suite.router.POST("/register", suite.userHandler.RegisterUser)
	suite.router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	// Unmarshal response body
	var respBody map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.NoError(suite.T(), err)

	// Assert response body
	assert.Equal(suite.T(), "Success created user", respBody["Message"])
	assert.Equal(suite.T(), &users[0], respBody["Data"])
}
