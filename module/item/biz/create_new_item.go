package biz

import (
	"context"
	"todolist/module/item/model"
)

// Handler -> Biz [-> Repository] -> Storage
/*
	Storage: Communicate with DB engine
	Repository:  Summarize and transform data to demanding struct for biz
	Biz: Use case, do business logic depending on requirements
	Handler: Send JSON to Clients
	These layers do not call each other directly, just by Interface
*/

/*
	In Go, Interface is declared at where it is used
*/

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
