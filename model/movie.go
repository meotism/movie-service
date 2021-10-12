package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID               int64          `gorm:"column:id;primary_key"  json:"id"`
	IsAdult          int            `gorm:"column:is_adult"  json:"is_adult"`
	Image            string         `gorm:"column:image"  json:"image"`
	OriginalLanguage string         `gorm:"column:original_language"  json:"original_language"`
	OriginalTitle    string         `gorm:"column:original_title"  json:"original_title"`
	Overview         string         `gorm:"column:overview"  json:"overview"`
	Popularity       float64        `gorm:"column:popularity"  json:"popularity"`
	MovieLink        string         `gorm:"column:movie_link"  json:"movie_link"`
	ReleaseDate      time.Time      `gorm:"column:release_date"  json:"release_date"`
	Duration         int            `gorm:"column:duration"  json:"duration"`
	SpokenLanguage   string         `gorm:"column:spoken_language"  json:"spoken_language"`
	RatingAverage    float64        `gorm:"column:rating_average"  json:"rating_average"`
	Status           int            `gorm:"column:status"  json:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}

type MovieFilter struct {
	IsAdult   int   `json:"is_adult"`  //default value no filter: -1
	ActorID   int64 `json:"actor_id"`	
	CateID    int64 `json:"cate_id"`
	MinRating int   `json:"min_rating"` //default value no filter: -1
}

type MovieResult struct {
	Data []Movie `json:"data"`
	Paginator
}
