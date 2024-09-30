package domain

import (
	"dev-oleksandrv/taskera-app/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	Spaces []Space `gorm:"many2many:space_users"`
	Lists  []List  `gorm:"foreignKey:CreatorID"`
}

func (user *User) BeforeCreate(_ *gorm.DB) error {
	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPassword
	}
	return nil
}
