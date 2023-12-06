package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	middleware "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := gin.Default()
	r.ContextWithFallback = true
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorMiddleware())

	handlers := InitHandlers()

	v1 := r.Group("/v1")
	v1.POST("/register", handlers.UserHandler.RegisterUser)
	v1.POST("/login", handlers.UserHandler.LoginUser)
	v1.POST("/forgot-password", handlers.UserHandler.ForgotPassword)
	v1.POST("/reset-password", handlers.UserHandler.ResetPassword)

	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":" + AppPort(),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutDuration()*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
