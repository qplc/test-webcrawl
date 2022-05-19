package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/test-webcrawl/rest"
	"bitbucket.org/test-webcrawl/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

var srv = &http.Server{}

func init() {

	//Initialize logger
	utils.InitializeLogger()

	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	router.Get("/health/shallow", rest.HealthCheckAPI)
	router.Post("/webcrawl", rest.WebCrawlAPI)

	appRunningPort := "8082"

	srv = &http.Server{
		Addr:         ":" + appRunningPort,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	utils.LogD.Info("Application started,", zap.Any("port", appRunningPort))
	srv.ListenAndServe()

	go GracefulShutdown()
}

//GracefulShutdown for graceful shutdown activities
func GracefulShutdown() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		utils.LogD.Info("Gracefully shutting down Analytics Producer", zap.Any("Signal", sig))
		//to stop server
		shutdownServer()

		done <- true
		utils.LogD.Info("Gracefully shut down Analytics Producer")
	}()

	<-waitForGracefulShutdown(done)
}

//Method for handling graceful server shutdown
func shutdownServer() {
	var duration = 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		utils.LogD.Error("Server Shutdown Failed", zap.Error(err))
	}

	utils.LogD.Info("Server Exited Properly")
}

//Method for handling graceful shutdown
func waitForGracefulShutdown(done chan bool) chan bool {
	return done
}

func main() {
	//Initialize logger
	utils.InitializeLogger()
}
