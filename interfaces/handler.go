package interfaces

import "github.com/gin-gonic/gin"

type (
	UserHandler interface {
		RegisterUser(c *gin.Context)
		LoginUser(c *gin.Context)
	}

	TransactionHandler interface {
		TopupTransaction(c *gin.Context)
		TransferTransaction(c *gin.Context)
	}
)
