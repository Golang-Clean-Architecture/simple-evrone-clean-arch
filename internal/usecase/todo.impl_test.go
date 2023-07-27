package usecase

import (
	"errors"
	"example-evrone/internal/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var todoRepository = &TodoRepositoryMock{Mock: mock.Mock{}}
var todoService = TodoServiceImpl{repo: todoRepository}

func TestMain(m *testing.M) {
	fmt.Println("Starting the Test")

	m.Run() // mengeksekusi semua unit test

	fmt.Println("Ending the test")
}

func TestTodoService_Get(t *testing.T) {
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
