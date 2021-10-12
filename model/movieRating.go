package model

type MovieRating struct {
	ID      int64 `gorm:"column:id;primary_key"  json:"id"`
	UserID  int64 `gorm:"column:user_id"  json:"user_id"`
	MovieID int64 `gorm:"column:movie_id"  json:"movie_id"`
	Rating  int `gorm:"column:rating"  json:"rating"`
}
