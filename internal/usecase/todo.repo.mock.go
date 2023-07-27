package usecase

import (
	"errors"
	"example-evrone/internal/entity"

	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TodoRepositoryMock) GetTodo(name *string) (*entity.Todo, error) {
	arguments := repository.Mock.Called(name)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found")
	} else {
		todo := arguments.Get(0).(entity.Todo)
		return &todo, nil
	}
}

func (repository *TodoRepositoryMock) GetAll() ([]*entity.Todo, error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("document is empty")
	} else {
		todo := arguments.Get(0).([]*entity.Todo)
		return todo, nil
	}
}

func (repository *TodoRepositoryMock) CreateTodo(entitys *entity.Todo) error {
	arguments := repository.Mock.Called(entitys)
	if arguments.Get(0).(entity.Todo).Name == "" {
		return errors.New("please fill todo name")
	} else {
		return nil
	}
}

func (repository *TodoRepositoryMock) DeleteTodo(entitys *string) error {
	arguments := repository.Mock.Called(entitys)
	if arguments.Get(0).(entity.Todo).Name != *entitys {
		return errors.New("no matched document found for delete")
	} else {
		return nil
	}
}

func (repository *TodoRepositoryMock) UpdateTodo(entitys *entity.Todo) error {
	arguments := repository.Mock.Called(entitys)
	if arguments.Get(0).(entity.Todo).Name != entitys.Name {
		return errors.New("no matched document found for update")
	} else {
		return nil
	}
}
