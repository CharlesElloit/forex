package models

import (
	"time"

	"github.com/google/uuid"
)

type ChargesType struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	Name             string    `json:"name" gorm:"type:varchar(100);not null"`
	Description      string    `json:"description" gorm:"type:varchar(200);null"`
	CreatedDate      string    `json:"created_date" gorm:"default:current_date"`
	CreatedTime      string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `json:"createdby_id" gorm:"type:uuid;null"`
	LastmodifiedDate string    `json:"lastmodified_date"`
	LastmodifiedTime time.Time `json:"lastmodified_time"`
	LastmodifiedbyID uuid.UUID `json:"lastmodifiedby_id" gorm:"type:uuid;null"`
	IsActive         bool      `json:"is_active" gorm:"type:boolean;default:false"`
	IsApproved       bool      `json:"is_approved" gorm:"type:boolean;default:false"`
	ApprovedTime     time.Time `json:"approved_time"`
	ApprovedbyID     uuid.UUID `json:"approvedby_id" gorm:"type:uuid;null"`
	SourceIp         string    `json:"source_ip" gorm:"type:varchar(200);not null"`
	SourceBrowser    string    `json:"source_browser" gorm:"type:varchar(200);not null"`
	RejectedDate     string    `json:"rejected_date" gorm:"null"`
	RejectedTime     time.Time `json:"rejected_time" gorm:"null"`
	RejectedbyID     uuid.UUID `json:"rejectedby_id" gorm:"type:uuid;null"`

	FkCreatedbyID      User `gorm:"foreignKey:CreatedbyID;"`
	FkApprovedbyID     User `gorm:"foreignKey:ApprovedbyID;"`
	FkRejectedbyID     User `gorm:"foreignKey:RejectedbyID;"`
	FkLastmodifiedbyID User `gorm:"foreignKey:LastmodifiedbyID;"`
}
