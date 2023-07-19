package usecase

import "example-evrone/internal/entity"

// here we create API/Service Contract
type TodoService interface {
	CreateTodo(*entity.Todo) error
	GetTodo(*string) (*entity.Todo, error)
	GetAll() ([]*entity.Todo, error)
	UpdateTodo(*entity.Todo) error
	DeleteTodo(*string) error
}
