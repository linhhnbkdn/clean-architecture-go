package biz

import (
	"cleanArchitechureGo/module/item/model"
	"context"
)

// Handler -> Biz [-> Repository] -> Storage
// Storage tầng thực hiện các thao tác với database như lưu, lấy dữ liệu, cập nhật dữ liệu
// Repository tầng tổng hợp dữ liệu và transform dữ liệu từ Storage trước khi trả về cho Biz hoặc ngược lại từ Biz trước khi lưu vào Storage
// Biz tầng chứa các logic của nghiệp vụ

type ICreateNewItemStore interface {
	Insert(ctx context.Context, item *model.ToDoItemDTOCreation) error
}

type createNewItemBiz struct {
	store ICreateNewItemStore
}

func NewCreateNewItemBiz(store ICreateNewItemStore) *createNewItemBiz {
	return &createNewItemBiz{store: store}
}

func (biz *createNewItemBiz) CreateNewItem(ctx context.Context, item *model.ToDoItemDTOCreation) error {
	if err := item.Validate(); err != nil {
		return err
	}
	if err := biz.store.Insert(ctx, item); err != nil {
		return err
	}
	return nil
}
