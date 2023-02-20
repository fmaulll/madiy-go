package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Point     int       `gorm:"default:0" json:"point"`
}
