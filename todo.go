package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" `
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ItemId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"Description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("i wanna sleep soo bad but update structure has no values ")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("i wanna sleep soo bad but update structure has no values ")
	}

	return nil
}
