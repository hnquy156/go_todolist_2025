package biz

import (
	"context"
	"todolist/common"
	"todolist/module/item/model"
)

type GetItemsStorage interface {
	GetItems(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.TodoItem, error)
}

type getItemsBiz struct {
	store GetItemsStorage
}

func NewGetItemsBiz(store GetItemsStorage) *getItemsBiz {
	return &getItemsBiz{store: store}
}

func (biz *getItemsBiz) GetItems(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.TodoItem, error) {
	data, err := biz.store.GetItems(ctx, paging, filter)
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}
