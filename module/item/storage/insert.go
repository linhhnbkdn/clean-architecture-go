package storage

import (
	"cleanArchitechureGo/module/item/model"
	"context"
)

func (s *SqlStore) Insert(ctx context.Context, item *model.ToDoItemDTOCreation) error {
	if err := item.Validate(); err != nil {
		return err
	}
	todoItem := &model.ToDoItem{
		Title:       item.Title,
		Description: item.Description,
		Status:      item.Status,
	}
	if err := s.db.Create(todoItem).Error; err != nil {
		return err
	}
	return nil
}
