package domain

import (
	"github.com/google/uuid"
	"time"
)

type Space struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Users       []User    `gorm:"many2many:space_users;constraint:OnDelete:CASCADE;"`
}

type SpaceWithRole struct {
	Space
	Role Role
}
