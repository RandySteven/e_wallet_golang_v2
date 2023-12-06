package middleware

import (
	"assignment_4/apperror"
	"assignment_4/entities/payload/res"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	var (
		errFieldValidation       *apperror.ErrFieldValidation
		errAmountLimit           *apperror.ErrAmountLimit
		errBalanceNotEnough      *apperror.ErrBalanceNotEnough
		errSenderAndReceiverSame *apperror.ErrSenderAndReceiverSame
		errDataNotFound          *apperror.ErrDataNotFound
		errWalletInvalid         *apperror.ErrWalletInvalid
		errInvalidRequest        *apperror.ErrInvalidRequest
		errTokenAlreadyUsed      *apperror.ErrTokenAlreadyUsed
		errTokenExpired          *apperror.ErrTokenExpired
		errTokenInvalid          *apperror.ErrTokenInvalid
		errInvalidFormat         *apperror.ErrInvalidFormat
		errLogin                 *apperror.ErrLogin
		errEmailAlreadyExists    *apperror.ErrEmailAlreadyExists
	)

	return func(c *gin.Context) {
		c.Next()

		resp := res.Response{}

		for _, ginErr := range c.Errors {
			resp.Errors = append(resp.Errors, ginErr.Err.Error())
			switch {
			case errors.As(ginErr.Err, &errSenderAndReceiverSame):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errAmountLimit):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errWalletInvalid):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errDataNotFound):
				c.AbortWithStatusJSON(http.StatusNotFound, resp)
			case errors.As(ginErr.Err, &errBalanceNotEnough):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errInvalidRequest):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errTokenAlreadyUsed):
				c.AbortWithStatusJSON(http.StatusForbidden, resp)
			case errors.As(ginErr.Err, &errTokenExpired):
				c.AbortWithStatusJSON(http.StatusForbidden, resp)
			case errors.As(ginErr.Err, &errTokenInvalid):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errFieldValidation):
				messages := strings.Split(ginErr.Err.Error(), "|")
				resp.Errors = messages
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errInvalidFormat):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errLogin):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errEmailAlreadyExists):
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": ginErr.Err.Error()})
			}
		}
	}
}
