package apps

import (
	middleware "assignment_4/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	userRouter := r.Group("users")
	userRouter.GET("/:id", h.UserHandler.GetUserById)

	r.Use(middleware.AuthMiddleware)

	r.GET("/hello", func(ctx *gin.Context) {
		time.Sleep(time.Second * 5)
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	transferRouter := r.Group("transfers")
	transferRouter.POST("", h.TransactionHandler.TransferTransaction)

	topupRouter := r.Group("topups")
	topupRouter.POST("", h.TransactionHandler.TopupTransaction)

	transactionsRouter := r.Group("transactions")
	transactionsRouter.GET("", h.TransactionHandler.GetAllTransactionsRecords)
	transactionsRouter.GET("histories", h.TransactionHandler.GetAllHistoryUserTransactions)

	gameRouter := r.Group("games")
	gameRouter.POST("", h.GameHandler.PlayGame)
	gameRouter.PUT("/:id", h.GameHandler.ChooseBox)
}
