package model

import "gorm.io/gorm"

const PageSize = 20

// Paginator .
type Paginator struct {
	Page  int   `json:"page,omitempty"`
	Limit int   `json:"limit,omitempty"`
	Total int64 `json:"total,omitempty"`
}

func (p Paginator) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.Limit

		return db.Offset(offset).Limit(p.Limit)
	}
}
