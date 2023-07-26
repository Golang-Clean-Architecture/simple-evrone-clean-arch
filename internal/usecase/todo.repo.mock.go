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
		return nil, errors.New("not found")
	} else {
		// todo := arguments.Get(0).(entity.Todo)
		return nil, nil
	}
}

func (repository *TodoRepositoryMock) CreateTodo(entity *entity.Todo) error {
	arguments := repository.Mock.Called(entity)
	if arguments.Get(0) == nil {
		return errors.New("not created")
	} else {
		return nil
	}
}

func (repository *TodoRepositoryMock) DeleteTodo(entity *string) error {
	arguments := repository.Mock.Called(entity)
	if arguments.Get(0) == nil {
		return errors.New("not created")
	} else {
		return nil
	}
}

func (repository *TodoRepositoryMock) UpdateTodo(entity *entity.Todo) error {
	arguments := repository.Mock.Called(entity)
	if arguments.Get(0) == nil {
		return errors.New("not created")
	} else {
		return nil
	}
}
