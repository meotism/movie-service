package model

type Producer struct {
	ID       int64    `gorm:"column:id;primary_key"  json:"id"`
	LogoPath string `gorm:"column:logo_path"  json:"logo_path"`
	Name     string `gorm:"column:name"  json:"name"`
}
