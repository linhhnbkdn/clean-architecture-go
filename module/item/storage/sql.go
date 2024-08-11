package storage

import "gorm.io/gorm"

type SqlStore struct {
	db *gorm.DB
}

func NewSQLStorage(db *gorm.DB) *SqlStore {
	return &SqlStore{db: db}
}
