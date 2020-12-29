package service_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yet-another-todo-list-golang/model/entity"
	"os"
	"testing"
	"time"

	mocks "github.com/yet-another-todo-list-golang/mocks/repository"
	"github.com/yet-another-todo-list-golang/service"
)

var todos []entity.Todo

func TestMain(m *testing.M) {
	for i := 1; i < 5; i++ {
		todos = append(todos, entity.Todo{
			Base: entity.Base{
				ID: fmt.Sprintf("35c54eae-ffaf-4085-a545-bb66d2abc22%d", i),
			},
			Name:        fmt.Sprintf("Todo%d", i),
			Description: fmt.Sprintf("Todo%d", i),
			Status:      "ACTIVE",
			UserID:      "",
			User:        entity.User{},
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		})
	}

	os.Exit(m.Run())
}

func TestTodoFindAll(t *testing.T) {
	todoRepository := new(mocks.TodoRepository)
	testCases := []struct {
		name   string
		entity []entity.Todo
		err    error
	}{
		{
			name:   "Simple read",
			entity: todos,
			err:    nil,
		},
		{
			name:   "Unknown error such as database error",
			entity: []entity.Todo{},
			err:    errors.New("database error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todoRepository.On("FindAll").Return(tc.err, tc.entity)
			todoService := service.NewTodoService(todoRepository)
			err, got := todoService.FindAll()
			if err != nil {
				assert.EqualError(t, err, "database error")
			}

			assert.Equal(t, got, todos)
		})
	}
}

func TestTodoFindOne(t *testing.T) {
}

func TestTodoCreate(t *testing.T) {

}

func TestTodoUpdate(t *testing.T) {

}

func TestTodoDelete(t *testing.T) {

}