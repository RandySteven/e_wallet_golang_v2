package apps

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	// r.Use(middleware.WithTimeOut)

	r.GET("/hello", func(ctx *gin.Context) {
		time.Sleep(time.Second * 5)
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

}
