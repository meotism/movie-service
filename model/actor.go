package model

import (
	"time"

	"gorm.io/gorm"
)

type Actor struct {
	ID           int64            `gorm:"column:id;primary_key"  json:"id"`
	Name         string         `gorm:"column:name"  json:"name"`
	Birthday     time.Time      `gorm:"column:birthday"  json:"birthday"`
	Deathday     time.Time      `gorm:"column:deathday"  json:"deathday"`
	Gender       int64            `gorm:"column:gender"  json:"gender"`
	PlaceOfBirth string         `gorm:"column:place_of_birth"  json:"place_of_birth"`
	Popularity   float64        `gorm:"column:popularity"  json:"popularity"`
	Avatar       string         `gorm:"column:avatar"  json:"avatar"`
	CreatedAt    time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}
