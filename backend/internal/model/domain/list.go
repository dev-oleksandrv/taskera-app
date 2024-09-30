package domain

import (
	"github.com/google/uuid"
	"time"
)

type List struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"default:null"`
	Emoji       string    `gorm:"size:255;not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	SpaceID   uuid.UUID `gorm:"type:uuid;not null"`
	Space     Space     `gorm:"foreignKey:SpaceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatorID uuid.UUID `gorm:"type:uuid;not null"`
	Creator   User      `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
