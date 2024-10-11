package domain

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Content     string    `gorm:"type:text;not null"`
	FullContent string    `gorm:"type:text;null"`
	Completed   bool      `gorm:"type:boolean;default:false"`
	Order       int       `gorm:"type:int;not null"`
	Deadline    time.Time `gorm:"type:timestamp;null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	AssigneeID *uuid.UUID `gorm:"type:uuid;null"`
	Assignee   User       `gorm:"foreignKey:AssigneeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ListID uuid.UUID `gorm:"type:uuid;not null"`
	List   List      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatorID uuid.UUID `gorm:"type:uuid;not null"`
	Creator   User      `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
