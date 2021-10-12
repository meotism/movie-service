package model

import "time"

type ViewHistorie struct {
	ID                int64     `gorm:"column:id;primary_key"  json:"id"`
	UserID            int64     `gorm:"column:user_id"  json:"user_id"`
	MovieID           int64     `gorm:"column:movie_id"  json:"movie_id"`
	LastMoviesRuntime int       `gorm:"column:last_movies_runtime"  json:"last_movies_runtime"`
	LastViewAt        time.Time `gorm:"column:last_view_at"  json:"last_view_at"`
}
