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
		// errWalletNumberInvalid   *apperror.ErrWalletNumberInvalid
		// errAmountInvalid         *apperror.ErrAmountInvalid
		errFieldValidation *apperror.ErrFieldValidation
		errAmountLimit     *apperror.ErrAmountLimit
		// errLengthValidation      *apperror.ErrLengthValidation
		errBalanceNotEnough      *apperror.ErrBalanceNotEnough
		errSenderAndReceiverSame *apperror.ErrSenderAndReceiverSame
		errDataNotFound          *apperror.ErrDataNotFound
		errWalletInvalid         *apperror.ErrWalletInvalid
		errInvalidRequest        *apperror.ErrInvalidRequest
		errTokenAlreadyUsed      *apperror.ErrTokenAlreadyUsed
		errTokenExpired          *apperror.ErrTokenExpired
		errTokenInvalid          *apperror.ErrTokenInvalid
	)

	return func(c *gin.Context) {
		c.Next()

		resp := res.Response{}

		for _, ginErr := range c.Errors {
			switch {
			case errors.As(ginErr.Err, &errSenderAndReceiverSame):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errAmountLimit):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errWalletInvalid):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errDataNotFound):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusNotFound, resp)
			case errors.As(ginErr.Err, &errBalanceNotEnough):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errInvalidRequest):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errTokenAlreadyUsed):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusForbidden, resp)
			case errors.As(ginErr.Err, &errTokenExpired):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusForbidden, resp)
			case errors.As(ginErr.Err, &errTokenInvalid):
				resp.Errors = append(resp.Errors, ginErr.Err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.As(ginErr.Err, &errFieldValidation):
				messages := strings.Split(ginErr.Err.Error(), "|")
				resp.Errors = messages
				c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.As(ginErr.Err, &errLengthValidation):
			// 	resp.Errors = append(resp.Errors, ginErr.Err.Error())
			// 	c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.As(ginErr.Err, &errFieldIsRequired):
			// 	resp.Errors = append(resp.Errors, ginErr.Err.Error())
			// 	c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": ginErr.Err.Error()})
			}
			return
		}
	}
}
