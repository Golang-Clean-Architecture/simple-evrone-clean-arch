package usecase

import (
	"example-evrone/internal/entity"
)

type TodoServiceImpl struct {
	repo TodoRepo
}

// make a function that act like a constructor
func NewTodoService(todoRepo TodoRepo) *TodoServiceImpl {
	return &TodoServiceImpl{
		repo: todoRepo,
	}
}

// receiver function or more like classes/struct method in python/java
func (u *TodoServiceImpl) CreateTodo(todo *entity.Todo) error {
	err := u.repo.CreateTodo(todo)
	return err
}

func (u *TodoServiceImpl) GetTodo(name *string) (*entity.Todo, error) {
	todo, err := u.repo.GetTodo(name)
	return todo, err
}

func (u *TodoServiceImpl) GetAll() ([]*entity.Todo, error) {
	todos, err := u.repo.GetAll()
	return todos, err
}

func (u *TodoServiceImpl) UpdateTodo(todo *entity.Todo) error {
	err := u.repo.UpdateTodo(todo)
	return err
}

func (u *TodoServiceImpl) DeleteTodo(name *string) error {
	err := u.repo.DeleteTodo(name)
	return err
}
