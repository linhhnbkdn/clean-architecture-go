package model

import (
	"cleanArchitechureGo/common"
	"errors"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
)

type ToDoItem struct {
	common.SQLModel
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (ToDoItem) TableName() string { return "todo_item" }

type ToDoItemDTOCreation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (ToDoItemDTOCreation) TableName() string { return "todo_item" }
func (t *ToDoItemDTOCreation) Validate() error {
	if t.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

type ToDoItemDTOUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (ToDoItemDTOUpdate) TableName() string { return "todo_item" }
