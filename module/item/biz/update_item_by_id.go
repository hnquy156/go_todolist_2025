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

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, data *model.TodoItemUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	item, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if item.Status == model.DeletedStatus {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
