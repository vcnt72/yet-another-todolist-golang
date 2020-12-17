package service

import (
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/repository"
)

// TodoService service of todo
type TodoService interface {
	FindAll() []entity.Todo
	Create(createDto dto.CreateTodoDto)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

// NewTodoService get service of todo
func NewTodoService() TodoService {
	todoRepository := repository.NewTodoRepository()
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (todoService *todoService) FindAll() []entity.Todo {
	return todoService.todoRepository.FindAll()
}

func (todoService *todoService) Create(createDto dto.CreateTodoDto) {
	todo := &entity.Todo{
		Name:        createDto.Name,
		Description: createDto.Description,
	}
	todoService.todoRepository.Create(*todo)
}