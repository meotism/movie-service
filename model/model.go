package model

import (
	"time"
)

// Condition Type.
const (
	ConditionTypeNot = "not"
	ConditionTypeOr  = "or"
)

type Model struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Condition .
type Condition struct {
	Type    string
	Pattern string
	Values  []interface{}
}
