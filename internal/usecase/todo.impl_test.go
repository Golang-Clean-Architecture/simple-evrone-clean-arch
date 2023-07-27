package usecase

import (
	"errors"
	"example-evrone/internal/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting the Test")

	m.Run() // mengeksekusi semua unit test

	fmt.Println("Ending the test")
}

func TestTodoService_Get(t *testing.T) {
	var todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
	var todoService = TodoServiceImpl{repo: todoRepository}
	t.Parallel()
	// program mock
	taskName := "Task 1"
	taskName2 := "Task 2"

	tests := []struct {
		name     string
		request  string
		expected *string
		err      error
		mock     func()
	}{
		{
			name:     "-------Retrieve correct TODO-------",
			request:  taskName,
			expected: &taskName,
			err:      nil,
			mock: func() {
				todoRepository.Mock.On("GetTodo", &taskName).Return(entity.Todo{Name: "Task 1", Status: "On Going"}, nil)
			},
		},
		{
			name:     "-------Retrieve false TODO-------",
			request:  taskName2,
			expected: nil,
			err:      errors.New("not found"),
			mock: func() {
				todoRepository.Mock.On("GetTodo", &taskName2).Return(nil, errors.New("not found"))
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			// t.Parallel()

			test.mock()
			todo, errors := todoService.GetTodo(&test.request)
			if todo == nil {
				assert.Nil(t, test.expected, todo)
				assert.Error(t, test.err, errors)
			} else {
				assert.Equal(t, test.expected, &todo.Name)
				assert.ErrorIs(t, test.err, errors)
			}
		})
	}
}

func TestTodoService_GetAll(t *testing.T) {
	t.Parallel()
	var todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
	var todoService = TodoServiceImpl{repo: todoRepository}
	// program mock
	sliceTodos := []*entity.Todo{
		{Name: "Task 1", Status: "On Going"},
		{Name: "Task 2", Status: "Complete"},
	}

	tests := []struct {
		name     string
		expected []*entity.Todo
		err      error
		mock     func()
	}{
		{
			name:     "-------Retrieve All TODO-------",
			expected: sliceTodos,
			err:      nil,
			mock: func() {
				todoRepository.Mock.On("GetAll").Return(sliceTodos, nil)
			},
		},
		{
			name:     "-------Retrieve None TODO-------",
			expected: nil,
			err:      errors.New("document is empty"),
			mock: func() {
				todoRepository.Mock.On("GetAll").Return(nil, errors.New("document is empty"))
			},
		},
	}

	for _, test := range tests {

		test := test
		todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
		todoService = TodoServiceImpl{repo: todoRepository}

		t.Run(test.name, func(t *testing.T) {
			test.mock()
			todos, errorNew := todoService.GetAll()
			if errorNew == nil {
				assert.ErrorIs(t, test.err, errorNew)
			} else {
				assert.Error(t, test.err, errorNew)
			}
			assert.Equal(t, test.expected, todos)
		})
	}
}

func TestTodoService_CreateTodo(t *testing.T) {
	t.Parallel()
	var todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
	var todoService = TodoServiceImpl{repo: todoRepository}
	// program mock
	entityNew := entity.Todo{Name: "Task 1", Status: "On Going"}
	entityFail := entity.Todo{Name: "", Status: "Done"}

	tests := []struct {
		name   string
		err    error
		mock   func()
		entity *entity.Todo
	}{
		{
			name:   "-------Create TODO-------",
			err:    nil,
			entity: &entityNew,
			mock: func() {
				todoRepository.Mock.On("CreateTodo", &entityNew).Return(entityNew)
			},
		},
		{
			name:   "-------Fail Create TODO-------",
			err:    errors.New("please fill todo name"),
			entity: &entityFail,
			mock: func() {
				todoRepository.Mock.On("CreateTodo", &entityFail).Return(entityFail)
			},
		},
	}

	for _, test := range tests {

		test := test
		todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
		todoService = TodoServiceImpl{repo: todoRepository}

		t.Run(test.name, func(t *testing.T) {
			test.mock()
			errorNew := todoService.CreateTodo(test.entity)
			if errorNew == nil {
				assert.ErrorIs(t, test.err, errorNew)
			} else {
				assert.Error(t, test.err, errorNew)
			}
		})
	}
}

func TestTodoService_UpdateTodo(t *testing.T) {
	t.Parallel()
	var todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
	var todoService = TodoServiceImpl{repo: todoRepository}
	// program mock
	entityNew := entity.Todo{Name: "Task 1", Status: "On Going"}
	entityFail := entity.Todo{Name: "", Status: "Done"}

	tests := []struct {
		name   string
		err    error
		mock   func()
		entity *entity.Todo
	}{
		{
			name:   "-------Update TODO-------",
			err:    nil,
			entity: &entityNew,
			mock: func() {
				todoRepository.Mock.On("UpdateTodo", &entityNew).Return(entityNew)
			},
		},
		{
			name:   "-------Fail Update TODO-------",
			err:    errors.New("please fill todo name"),
			entity: &entityFail,
			mock: func() {
				todoRepository.Mock.On("UpdateTodo", &entityFail).Return(entityNew)
			},
		},
	}

	for _, test := range tests {

		test := test
		todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
		todoService = TodoServiceImpl{repo: todoRepository}

		t.Run(test.name, func(t *testing.T) {
			test.mock()
			errorNew := todoService.UpdateTodo(test.entity)
			if errorNew == nil {
				assert.ErrorIs(t, test.err, errorNew)
			} else {
				assert.Error(t, test.err, errorNew)
			}
		})
	}
}

func TestTodoService_DeleteTodo(t *testing.T) {
	t.Parallel()
	var todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
	var todoService = TodoServiceImpl{repo: todoRepository}
	// program mock
	entityNew := entity.Todo{Name: "Task 1", Status: "On Going"}
	entityFail := entity.Todo{Name: "", Status: "Done"}

	tests := []struct {
		name   string
		err    error
		mock   func()
		entity *entity.Todo
	}{
		{
			name:   "-------Delete TODO-------",
			err:    nil,
			entity: &entityNew,
			mock: func() {
				todoRepository.Mock.On("DeleteTodo", &entityNew.Name).Return(entityNew)
			},
		},
		{
			name:   "-------Fail Delete TODO-------",
			err:    errors.New("please fill todo name"),
			entity: &entityFail,
			mock: func() {
				todoRepository.Mock.On("DeleteTodo", &entityFail.Name).Return(entityNew)
			},
		},
	}

	for _, test := range tests {

		test := test
		todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
		todoService = TodoServiceImpl{repo: todoRepository}

		t.Run(test.name, func(t *testing.T) {
			test.mock()
			errorNew := todoService.DeleteTodo(&test.entity.Name)
			if errorNew == nil {
				assert.ErrorIs(t, test.err, errorNew)
			} else {
				assert.Error(t, test.err, errorNew)
			}
		})
	}
}
