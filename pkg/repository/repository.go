package repository

import (
	rest_api "api-learn/rest-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rest_api.User) (int, error)
	GetUser(username string, password string) (rest_api.User, error)
}

type TodoList interface {
	Create(userId int, list rest_api.TodoList) (int, error)
	GetAll(userId int) ([]rest_api.TodoList, error)
	GetById(userId int, listId int) (rest_api.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input rest_api.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item rest_api.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]rest_api.TodoItem, error)
	GetById(userId int, itemId int) (rest_api.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input rest_api.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
