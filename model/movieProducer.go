package model

type MovieProducer struct {
	ID         int64 `gorm:"column:id;primary_key"  json:"id"`
	MovieID    int64 `gorm:"column:movie_id"  json:"movie_id"`
	ProducerID int64 `gorm:"column:producer_id"  json:"producer_id"`
}
