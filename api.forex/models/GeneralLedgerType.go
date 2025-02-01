package models

import (
	"time"

	"github.com/google/uuid"
)

type GlType struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	Name             string    `json:"name"`
	ShortName        string    `json:"short_name"`
	Description      string    `json:"description"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:varchar(15);not null;default:current_date"`
	CreatedTime      time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `json:"createdby_id" gorm:"type:uuid;foreignKey:UserID"`
	LastmodifiedDate string    `json:"lastmodified_date"`
	LastmodifiedTime time.Time `json:"lastmodified_time"`
	LastmodifiedbyID uuid.UUID `json:"lastmodifiedby_id" gorm:"type:uuid;foreignKey:UserID;default:null"`
	IsApproved       *bool     `json:"is_approved"`
	ApprovedByID     *string   `json:"approvedby_id"`
	UserID           uuid.UUID
}
