package models

import (
	"time"

	"github.com/google/uuid"
)

type GlGroup struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	GroupName        string    `json:"group_name" gorm:"type:varchar(100);not null"`
	GroupCode        uint      `json:"group_code" gorm:"type:int;not null"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime      string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `sjon:"createdby_id"`
	LastmodifiedDate string    `json:"lastmodified_date"`
	LastmodifiedTime string    `json:"lastmodified_time" gorm:"type:varchar(8);default:to_char(now(), 'HH24:MI:SS AM')"`
	LastmodifiedbyID uuid.UUID `json:"lastmodifiedby_id"`
	IsApproved       bool      `json:"is_approved"`
	ApprovedbyID     string    `json:"approvedby_id"`
	ApprovedTime     time.Time `json:"approved_time"`
	SourceIp         string    `json:"source_ip"`
	SourceBrowser    string    `json:"source_browser"`
	RejectedDate     time.Time `json:"rejected_date" gorm:"type:date;null"`
	RejectedTime     time.Time `json:"rejected_time"`
	RejectedbyID     uuid.UUID `json:"rejectedby_id"`

	FkCreatedbyID      User `gorm:"foreignKey:CreatedbyID;"`
	FkApprovedbyID     User `gorm:"foreignKey:ApprovedbyID;"`
	FkRejectedbyID     User `gorm:"foreignKey:RejectedbyID;"`
	FkLastmodifiedbyID User `gorm:"foreignKey:LastmodifiedbyID;"`
}
