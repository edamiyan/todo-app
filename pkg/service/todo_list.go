package service

import (
	"github.com/edamiyan/todo-app"
	"github.com/edamiyan/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, ListId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, ListId)
}

func (s *TodoListService) DeleteById(userId, ListId int) error {
	return s.repo.DeleteById(userId, ListId)
}

func (s *TodoListService) UpdateById(userId, ListId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(userId, ListId, input)
}
