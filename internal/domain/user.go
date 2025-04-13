package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email        string     `gorm:"unique;not null" json:"email"`
	PasswordHash string     `gorm:"not null" json:"-"`
	DisplayName  string     `gorm:"not null" json:"displayName"`
	Bio          string     `json:"bio"`
	AvatarURL    string     `json:"avatarUrl"`
	Preferences  []byte     `gorm:"type:jsonb" json:"preferences"`
	CreatedAt    time.Time  `gorm:"not null;default:now()" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"not null;default:now()" json:"updatedAt"`
	LastLoginAt  *time.Time `json:"lastLoginAt"`
}
