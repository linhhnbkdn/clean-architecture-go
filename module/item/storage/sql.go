package storage

import (
	"github.com/samber/do"
	"gorm.io/gorm"
)

type SqlStore struct {
	db *gorm.DB
}

func NewSQLStorage(i *do.Injector) *SqlStore {
	var database = do.MustInvoke[*gorm.DB](i)
	return &SqlStore{db: database}
}
