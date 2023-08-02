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
		assert   func(expected *string, reality *entity.Todo, expectedError error, realityError error)
	}{
		{
			name:     "-------Retrieve correct TODO-------",
			request:  taskName,
			expected: &taskName,
			err:      nil,
			mock: func() {
				todoRepository.Mock.On("GetTodo", &taskName).Return(entity.Todo{Name: "Task 1", Status: "On Going"}, nil)
			},
			assert: func(expected *string, reality *entity.Todo, expectedError error, realityError error) {
				assert.Equal(t, expected, &reality.Name)
				assert.ErrorIs(t, expectedError, realityError)
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
			assert: func(expected *string, reality *entity.Todo, expectedError, realityError error) {
				assert.Nil(t, expected, reality)
				assert.Error(t, expectedError, realityError)
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			// t.Parallel()

			test.mock()
			todo, errors := todoService.GetTodo(&test.request)

			fmt.Println("test 1")
			test.assert(test.expected, todo, test.err, errors)
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
		assert   func(expected []*entity.Todo, reality []*entity.Todo, expectedError error, realityError error)
	}{
		{
			name:     "-------Retrieve All TODO-------",
			expected: sliceTodos,
			err:      nil,
			mock: func() {
				todoRepository.Mock.On("GetAll").Return(sliceTodos, nil)
			},
			assert: func(expected, reality []*entity.Todo, expectedError, realityError error) {
				assert.ErrorIs(t, expectedError, realityError)
				assert.Equal(t, expected, reality)
			},
		},
		{
			name:     "-------Retrieve None TODO-------",
			expected: nil,
			err:      errors.New("document is empty"),
			mock: func() {
				todoRepository.Mock.On("GetAll").Return(nil, errors.New("document is empty"))
			},
			assert: func(expected, reality []*entity.Todo, expectedError, realityError error) {
				assert.Error(t, expectedError, realityError)
				assert.Equal(t, expected, reality)
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
			test.assert(test.expected, todos, test.err, errorNew)
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
		assert func(expectedError error, realityError error)
	}{
		{
			name:   "-------Create TODO-------",
			err:    nil,
			entity: &entityNew,
			mock: func() {
				todoRepository.Mock.On("CreateTodo", &entityNew).Return(entityNew)
			},
			assert: func(expectedError, realityError error) {
				assert.ErrorIs(t, expectedError, realityError)
			},
		},
		{
			name:   "-------Fail Create TODO-------",
			err:    errors.New("please fill todo name"),
			entity: &entityFail,
			mock: func() {
				todoRepository.Mock.On("CreateTodo", &entityFail).Return(entityFail)
			},
			assert: func(expectedError, realityError error) {
				assert.Error(t, expectedError, realityError)
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
			test.assert(test.err, errorNew)
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
		assert func(expectedError error, realityError error)
	}{
		{
			name:   "-------Update TODO-------",
			err:    nil,
			entity: &entityNew,
			mock: func() {
				todoRepository.Mock.On("UpdateTodo", &entityNew).Return(entityNew)
			},
			assert: func(expectedError, realityError error) {
				assert.ErrorIs(t, expectedError, realityError)
			},
		},
		{
			name:   "-------Fail Update TODO-------",
			err:    errors.New("please fill todo name"),
			entity: &entityFail,
			mock: func() {
				todoRepository.Mock.On("UpdateTodo", &entityFail).Return(entityNew)
			},
			assert: func(expectedError, realityError error) {
				assert.Error(t, expectedError, realityError)
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
			test.assert(test.err, errorNew)
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
		assert func(expectedError error, realityError error)
	}{
		{
			name:   "-------Delete TODO-------",
			err:    nil,
			entity: &entityNew,
			mock: func() {
				todoRepository.Mock.On("DeleteTodo", &entityNew.Name).Return(entityNew)
			},
			assert: func(expectedError, realityError error) {
				assert.ErrorIs(t, expectedError, realityError)
			},
		},
		{
			name:   "-------Fail Delete TODO-------",
			err:    errors.New("please fill todo name"),
			entity: &entityFail,
			mock: func() {
				todoRepository.Mock.On("DeleteTodo", &entityFail.Name).Return(entityNew)
			},
			assert: func(expectedError, realityError error) {
				assert.Error(t, expectedError, realityError)
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
			test.assert(test.err, errorNew)
		})
	}
}
