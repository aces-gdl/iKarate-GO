package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string
	Currency    string
	Amount      float64   `gorm:"type:numeric(8,2)"`
	SellUserID  uuid.UUID `gorm:"type:uuid;default:null;"`
	BuyUserID   uuid.UUID `gorm:"type:uuid;default:null;"`
	SellID      uuid.UUID `gorm:"type:uuid;default:null;"`
	Active      bool
}
