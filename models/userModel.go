package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Email        string         `gorm:"unique"`
	Password     string
	GoogleID     string
	ImageURL     string
	Name         string
	FamilyName   string
	GivenName    string
	PermissionID uuid.UUID
	Permissions  Permission `gorm:"foreignKey:PermissionID;references:ID"`
	Ranking      int
	CategoryID   uuid.UUID
	Category     Category `gorm:"foreignKey:CategoryID;references:ID"`
}
