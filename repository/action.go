package repository

import "gorm.io/gorm"

type IAction interface {
	authRepository
}

type store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *store {
	return &store{
		db: db,
	}
}
