package storage

import (
	"context"
	"todolist/common"
	"todolist/module/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
