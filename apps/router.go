package apps

import (
	middleware "assignment_4/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	r.Use(middleware.AuthMiddleware)

	r.GET("/hello", func(ctx *gin.Context) {
		time.Sleep(time.Second * 5)
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	transferRouter := r.Group("transfer")
	transferRouter.POST("", h.TransactionHandler.TransferTransaction)

	topupRouter := r.Group("topup")
	topupRouter.POST("", h.TransactionHandler.TopupTransaction)

}
