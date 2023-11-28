package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventsEnrollment struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string
	Price
	Address   string
	Phone     string
	email     string
	ManagerID uuid.UUID `gorm:"type:uuid;default:null;"`
	Manager   User      `gorm:"foreignKey:ManagerID;"`
	Active    bool
}
