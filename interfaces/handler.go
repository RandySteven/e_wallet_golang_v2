package interfaces

import "github.com/gin-gonic/gin"

type (
	UserHandler interface {
		RegisterUser(c *gin.Context)
		LoginUser(c *gin.Context)
		GetUserById(c *gin.Context)
		ForgotPassword(c *gin.Context)
		ResetPassword(c *gin.Context)
	}

	TransactionHandler interface {
		TopupTransaction(c *gin.Context)
		TransferTransaction(c *gin.Context)
		// GetAllHistoryUserTransactions(c *gin.Context)
		GetAllTransactionsRecords(c *gin.Context)
	}

	GameHandler interface {
		PlayGame(c *gin.Context)
		ChooseBox(c *gin.Context)
	}
)
