package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	CustomerId uuid.UUID `gorm:"type:uuid" json:"customerId"`
	UserId     uuid.UUID `gorm:"type:uuid" json:"userId"`
	Price      int64     `json:"price"`
}
