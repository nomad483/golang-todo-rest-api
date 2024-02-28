package repository

import (
	rest_api "api-learn/rest-api"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item rest_api.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId int, listId int) ([]rest_api.TodoItem, error) {
	var items []rest_api.TodoItem
	query := fmt.Sprintf(
		"SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId int, itemId int) (rest_api.TodoItem, error) {
	var item rest_api.TodoItem
	query := fmt.Sprintf(
		"SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2",
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *TodoItemPostgres) Update(userId int, itemId int, input rest_api.UpdateItemInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValue = append(setValue, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s li, %s ul WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d",
		todoListsTable,
		setQuery,
		listsItemsTable,
		usersListsTable,
		argId,
		argId+1,
	)

	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)

	return err
}
