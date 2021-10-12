package model

type MovieActor struct {
	ID            int64    `gorm:"column:id;primary_key"  json:"id"`
	MovieID       int64    `gorm:"column:movie_id"  json:"movie_id"`
	ActorID       int64    `gorm:"column:actor_id"  json:"actor_id"`
	CharacterName string `gorm:"column:character_name"  json:"character_name"`
	IsMainActor   int    `gorm:"column:is_main_actor"  json:"is_main_actor"`
}
