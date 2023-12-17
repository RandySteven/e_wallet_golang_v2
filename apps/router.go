package apps

import (
	middleware "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	r.Use(middleware.AuthMiddleware)
	userRouter := r.Group("users")
	userRouter.GET("", h.UserHandler.GetUserById)

	transferRouter := r.Group("transfers")
	transferRouter.POST("", h.TransactionHandler.TransferTransaction)

	topupRouter := r.Group("topups")
	topupRouter.POST("", h.TransactionHandler.TopupTransaction)

	transactionsRouter := r.Group("transactions")
	transactionsRouter.GET("", h.TransactionHandler.GetAllTransactionsRecords)

	gameRouter := r.Group("games")
	gameRouter.POST("", h.GameHandler.PlayGame)
	gameRouter.PUT("/:id", h.GameHandler.ChooseBox)
	gameRouter.GET("/chances", h.GameHandler.CurrentUserChance)
}
