package controller

import (
	"example-evrone/internal/entity"
	"example-evrone/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoUsecase usecase.TodoService
}

func NewTodoController(todoService usecase.TodoService) TodoController {
	return TodoController{
		TodoUsecase: todoService,
	}
}

func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todo entity.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.TodoUsecase.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TodoController) GetTodo(ctx *gin.Context) {
	var taskName string = ctx.Param("name")
	todo, err := tc.TodoUsecase.GetTodo(&taskName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (tc *TodoController) GetAll(ctx *gin.Context) {
	todos, err := tc.TodoUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (tc *TodoController) UpdateTodo(ctx *gin.Context) {
	var todo entity.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.TodoUsecase.UpdateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TodoController) DeleteTodo(ctx *gin.Context) {
	taskName := ctx.Param("name")
	err := tc.TodoUsecase.DeleteTodo(&taskName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// grouping the route for anything that has prefix /todo routes
func (tc *TodoController) RegisterTodoRoutes(registerGroup *gin.RouterGroup) {
	todoRoute := registerGroup.Group("/todo")

	todoRoute.POST("/create", tc.CreateTodo)
	todoRoute.GET("/get/:name", tc.GetTodo)
	todoRoute.GET("/get", tc.GetAll)
	todoRoute.POST("/update", tc.UpdateTodo)
	todoRoute.DELETE("/delete/:name", tc.DeleteTodo)
}
