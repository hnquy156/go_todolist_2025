package storage

import (
	"context"
	"todolist/common"
	"todolist/module/item/model"
)

func (s *sqlStore) GetItems(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.TodoItem, error) {
	var data []model.TodoItem

	db := s.db.Where("status <> ?", model.DeletedStatus)

	if filter != nil {
		db = db.Where(filter)
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Order(" id desc").
		Offset(offset).
		Limit(paging.Limit).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
