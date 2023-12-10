package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketDetail struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string
	SellUserID  uuid.UUID `gorm:"type:uuid;default:null;"`
	BuyUserID   uuid.UUID `gorm:"type:uuid;default:null;"`
	TotalPaid   float64   `gorm:"type:numeric(8,2)"`
	TicketID    uuid.UUID
	Active      bool
}
