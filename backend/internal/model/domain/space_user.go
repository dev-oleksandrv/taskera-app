package domain

import (
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	Owner  Role = "OWNER"
	Admin  Role = "ADMIN"
	Editor Role = "EDITOR"
)

type SpaceUser struct {
	SpaceID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Role      Role
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
