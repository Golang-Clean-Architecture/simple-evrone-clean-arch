// Package app configures and runs application.
package app

import (
	"context"
	"example-evrone/config"
	"example-evrone/internal/controller"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/internal/usecase/repo"
	"github.com/evrone/go-clean-template/internal/usecase/webapi"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	defer mongoClient.Disconnect(ctx)

	// Repository
	ctx := context.TODO()
	mongoConn := options.Client().ApplyURI(cfg.Mongo.URL)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection is successfull!")
	// Use case
	translationUseCase := usecase.New(
		repo.New(mongoClient),
		webapi.New(),
	)

	// HTTP Server
	// handler := gin.New()
	// v1.NewRouter(handler, l, translationUseCase)
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	var server *gin.Engine = gin.Default
	basePath := server.Group("/v1")
	controller.UserController.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":8080"))

	// // Waiting signal
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// select {
	// case s := <-interrupt:
	// 	l.Info("app - Run - signal: " + s.String())
	// case err = <-httpServer.Notify():
	// 	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

	// // Shutdown
	// err = httpServer.Shutdown()
	// if err != nil {
	// 	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	// }
}
