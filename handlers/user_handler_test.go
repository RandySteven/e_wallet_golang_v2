package handlers_test

import (
	"assignment_4/apperror"
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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (suite *UserHandlerTestSuite) TestGetUserByInvalidId() {

	req, _ := http.NewRequest("GET", "/users/A", nil)
	w := httptest.NewRecorder()

	suite.router.GET("/users/:id", suite.userHandler.GetUserById)
	suite.router.ServeHTTP(w, req)
	suite.T().Log("response : ", w.Body)

	suite.Assert().Equal(http.StatusBadRequest, w.Code)
}

func (suite *UserHandlerTestSuite) TestGetUserNotFound() {
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	suite.userUsecase.On("GetUserDetail", mock.Anything, uint(1)).Return(nil, &apperror.ErrDataNotFound{Data: "user"})
	suite.router.GET("/users/:id", suite.userHandler.GetUserById)
	suite.router.ServeHTTP(w, req)
	suite.T().Log("response : ", w.Body)

	suite.Assert().Equal(http.StatusNotFound, w.Code)
}

func (suite *UserHandlerTestSuite) TestGetUserByIdInternalServerError() {

	suite.userUsecase.On("GetUserDetail", mock.Anything, uint(1)).Return(nil, errors.New("mock error"))

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	suite.router.GET("/users/:id", suite.userHandler.GetUserById)
	suite.router.ServeHTTP(w, req)
	suite.T().Log("response : ", w.Body)

	suite.Assert().Equal(http.StatusInternalServerError, w.Code)
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

func (suite *UserHandlerTestSuite) TestLoginUserFailedInternalServerError() {
	loginRequest := &req.UserLoginRequest{
		Email:    "randy.steven@gmail.com",
		Password: "test_1234",
	}

	suite.userUsecase.On("LoginUser", mock.Anything, mock.AnythingOfType("*req.UserLoginRequest")).Return(nil, errors.New("mock error"))

	loginRequestBody, err := json.Marshal(loginRequest)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginRequestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()

	suite.router.POST("/login", suite.userHandler.LoginUser)
	suite.router.ServeHTTP(w, req)

	suite.T().Log(w.Body)
	suite.Equal(http.StatusInternalServerError, w.Code)

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

//Test Forgot Password
func (suite *UserHandlerTestSuite) TestForgotPassword() {
	forgotPassword := &req.ForgotPasswordRequest{
		Email: "randy.steven@gmail.com",
	}

	forgotPasswordToken := &models.ForgotPasswordToken{
		ResetToken:  uuid.NewString(),
		TokenExpiry: time.Now().Add(time.Minute * 15),
		IsValid:     true,
	}

	suite.userUsecase.
		On("ForgotPassword", mock.Anything, mock.AnythingOfType("*req.ForgotPasswordRequest")).
		Return(forgotPasswordToken, nil)

	forgotPasswordRequestBody, _ := json.Marshal(forgotPassword)
	req, _ := http.NewRequest("POST", "/forgot-password", bytes.NewBuffer(forgotPasswordRequestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/forgot-password", suite.userHandler.ForgotPassword)
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *UserHandlerTestSuite) TestForgotPasswordBadRequest() {
	forgotPassword := &req.ForgotPasswordRequest{
		Email: "randy.steven",
	}

	forgotPasswordRequestBody, _ := json.Marshal(forgotPassword)
	req, _ := http.NewRequest("POST", "/forgot-password", bytes.NewBuffer(forgotPasswordRequestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/forgot-password", suite.userHandler.ForgotPassword)
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *UserHandlerTestSuite) TestForgotPasswordInternalServerError() {
	forgotPassword := &req.ForgotPasswordRequest{
		Email: "randy.steven@gmail.com",
	}

	suite.userUsecase.
		On("ForgotPassword", mock.Anything, mock.AnythingOfType("*req.ForgotPasswordRequest")).
		Return(nil, errors.New("mock error"))

	forgotPasswordRequestBody, _ := json.Marshal(forgotPassword)
	req, _ := http.NewRequest("POST", "/forgot-password", bytes.NewBuffer(forgotPasswordRequestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/forgot-password", suite.userHandler.ForgotPassword)
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *UserHandlerTestSuite) TestResetPassword() {
	resetPassword := &req.PasswordResetRequest{
		Email:       "randy.steven@gmail.com",
		NewPassword: "test_1234",
	}

	suite.userUsecase.
		On("ResetPassword", mock.Anything, mock.AnythingOfType("*req.PasswordResetRequest")).
		Return(&users[0], nil)

	resetPasswordRequest, _ := json.Marshal(resetPassword)
	req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(resetPasswordRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/reset-password", suite.userHandler.ResetPassword)
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *UserHandlerTestSuite) TestResetPasswordBadRequest() {
	resetPassword := &req.PasswordResetRequest{
		Email: "randy.steven@gmail.com",
	}

	resetPasswordRequest, _ := json.Marshal(resetPassword)
	req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(resetPasswordRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/reset-password", suite.userHandler.ResetPassword)
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *UserHandlerTestSuite) TestResetPasswordInternalServerError() {
	resetPassword := &req.PasswordResetRequest{
		Email:       "randy.steven@gmail.com",
		NewPassword: "test_1234",
	}

	suite.userUsecase.
		On("ResetPassword", mock.Anything, mock.AnythingOfType("*req.PasswordResetRequest")).
		Return(nil, errors.New("mock error"))

	resetPasswordRequest, _ := json.Marshal(resetPassword)
	req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(resetPasswordRequest))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.POST("/reset-password", suite.userHandler.ResetPassword)
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusInternalServerError, w.Code)
}
