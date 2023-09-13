package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   uuid.UUID
	Name string
	gorm.Model
}
