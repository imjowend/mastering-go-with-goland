package todo

import (
	"errors"
	"strings"
)

type Service struct {
	todos []Item
}

type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (svc *Service) Search(query string) []string {
	var results []string
	for _, todo := range svc.todos {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			results = append(results, todo.Task)
		}
	}
	return results
}

func (svc *Service) Add(todo string) error {
	for _, t := range svc.todos {
		if t.Task == todo {
			return errors.New("todo already exists")
		}
	}
	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})
	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todos
}
