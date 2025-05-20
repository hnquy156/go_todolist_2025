package model

import (
	"errors"
	"strings"
	"todolist/common"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title can not be empty")
	ErrItemIsDeleted      = errors.New("this item is deleted")
	DeletedStatus         = "Deleted"
	EntityName            = "Item"
)

// Tag json, khi giao tiep vs API se gtiep client thong qua ngon ngu trung gian chinh la javascript object notation
type TodoItem struct {
	common.SQLModel
	Title       string        `json:"title" gorm:"column:title"`
	Description string        `json:"description" gorm:"column:description"`
	Status      string        `json:"status" gorm:"column:status"`
	Image       *common.Image `json:"image" gorm:"column:image"`
}

type TodoItemCreation struct {
	Id          int           `json:"-" gorm:"column:id"`
	Title       string        `json:"title" gorm:"column:title"`
	Description string        `json:"description" gorm:"column:description"`
	Image       *common.Image `json:"image" gorm:"column:image"`
	// Status      string `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return "todo_items"
}

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)
	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func (TodoItemUpdate) TableName() string {
	return "todo_items"
}

func (i *TodoItemUpdate) Validate() error {
	i.Title = strings.TrimSpace(i.Title)
	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}
