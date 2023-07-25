// Package app configures and runs application.
package app

import (
	"context"
	"example-evrone/config"
	"example-evrone/internal/controller"
	"example-evrone/internal/usecase"
	"example-evrone/internal/usecase/repo"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-gonic/gin"
)

var (
	server         *gin.Engine
	todoService    usecase.TodoServiceImpl
	todoController controller.TodoController
	ctx            context.Context
	todoCollection *mongo.Collection
	mongoClient    *mongo.Client
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
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
	todoCollection = mongoClient.Database("tododb").Collection("todos")
	// Use case
	todoService = *usecase.NewTodoService(repo.New(todoCollection, ctx))

	// // Controller
	todoController = controller.NewTodoController(todoService)

	// HTTP Server
	server = gin.Default()
	basePath := server.Group("/v1")
	todoController.RegisterTodoRoutes(basePath, todoService)

	log.Fatal(server.Run(":8080"))

}
