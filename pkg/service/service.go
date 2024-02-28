package service

import (
	rest_api "api-learn/rest-api"
	"api-learn/rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list rest_api.TodoList) (int, error)
	GetAll(userId int) ([]rest_api.TodoList, error)
	GetById(userId int, listId int) (rest_api.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input rest_api.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item rest_api.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]rest_api.TodoItem, error)
	GetById(userId int, itemId int) (rest_api.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input rest_api.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
