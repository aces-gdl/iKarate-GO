package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dojo struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	ShortName   string
	Address     string
	Phone       string
	Email       string
	ManagerID   uuid.UUID `gorm:"type:uuid;default:null;"`
	Manager     User      `gorm:"foreignKey:ManagerID;"`
	Active      bool
}
