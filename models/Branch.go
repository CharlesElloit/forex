package models

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	Name             string    `json:"name" gorm:"type:varchar(30)"`
	Code             int       `json:"code" gorm:"type:int;not null"`
	OpenDate         string    `json:"open_date" gorm:"type:date;default:current_date"`
	CloseDate        string    `json:"close_date" gorm:"type:date;null"`
	Address          string    `json:"address" gorm:"type:varchar(200);null"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Type             string    `json:"type"`
	Contact          string    `json:"contact" gorm:"type:varchar(14)"`
	EmailAddres      string    `json:"email_address"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:varchar(15);not null;default:current_date"`
	CreatedTime      time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `gorm:"not null;index"`
	LastmodifiedDate string    `json:"lastmodified_date" gorm:"type:date;null"`
	LastmodifiedTime time.Time `json:"lastmodified_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	LastmodifiedbyID uuid.UUID `gorm:"not null;index"`
	ApprovedDate     time.Time `json:"approved_date" gorm:"type:date;null"`
	ApprovedTime     time.Time `json:"approved_time"`
	IsApproved       bool      `json:"is_approved"`
	ApprovedbyID     uuid.UUID `json:"approvedby_id"`
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
