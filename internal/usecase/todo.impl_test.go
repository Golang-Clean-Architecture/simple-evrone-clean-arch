package usecase

import (
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
	// t.Parallel()
	// program mock
	taskName := "Task 1"
	tests := []struct {
		name     string
		request  string
		expected *string
		mock     func(string)
	}{
		{
			name:     "-------Retrieve correct TODO-------",
			request:  taskName,
			expected: &taskName,
			mock: func(str string) {
				todoRepository.Mock.On("GetTodo", &str).Return(entity.Todo{Name: "Task 1", Status: "On Going"})
			},
		},
		{
			name:     "-------Retrieve false TODO-------",
			request:  taskName,
			expected: &taskName,
			mock: func(str string) {
				todoRepository.Mock.On("GetTodo", &str).Return(nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// test := test
			fmt.Println(test.name)
			test.mock(test.request)
			todo, _ := todoService.GetTodo(&test.request)
			fmt.Println(todo)
			assert.Equal(t, &todo.Name, test.expected)
		})
	}
}
