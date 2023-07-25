package controller

import (
	"example-evrone/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoUsecase usecase.TodoServiceImpl
}

// grouping the route for anything that has prefix /todo routes
func (tc *TodoController) RegisterTodoRoutes(registerGroup *gin.RouterGroup, service usecase.TodoServiceImpl) {
	todoRoute := registerGroup.Group("/todo")
	tc.TodoUsecase = service

	todoRoute.POST("/create", tc.CreateTodo)
	todoRoute.GET("/get/:name", tc.GetTodo)
	todoRoute.GET("/get", tc.GetAll)
	todoRoute.POST("/update", tc.UpdateTodo)
	todoRoute.DELETE("/delete/:name", tc.DeleteTodo)
}
