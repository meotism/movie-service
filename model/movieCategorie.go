package model

type MovieCategorie struct {
	ID      int64 `gorm:"column:id;primary_key"  json:"id"`
	MovieID int64 `gorm:"column:movie_id"  json:"movie_id"`
	CateID  int64 `gorm:"column:cate_id"  json:"cate_id"`
}
