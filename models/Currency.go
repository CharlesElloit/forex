package models

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	ID                       uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	CurrencyName             string    `json:"currency_name" gorm:"unique"`
	CurrencyCode             string    `json:"currency_code" gorm:"unique"`
	DecimalPlaces            int       `json:"decimal_places"`
	DisplaySymbol            string    `json:"display_symbol"`
	CurrencyStatus           string    `json:"currency_status"`
	InternationlizedNameCode string    `json:"internationalized_name_code"`
	ActivitionDate           string    `json:"activition_date"`
	InactiveDate             string    `json:"inactive_date"`
	CreatedDate              time.Time `json:"created_date" gorm:"default:current_date"`
	CreatedTime              string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID              uuid.UUID `json:"createdby_id"`
	LastmodifiedDate         time.Time `json:"lastmodified_date"`
	LastmodifiedbyID         uuid.UUID `json:"lastmodifiedby_id"`
	IsApproved               bool      `json:"is_approved" gorm:"default:false"`
	ApprovedbyID             uuid.UUID `json:"approvedby_id"`
	ApprovedTime             time.Time `json:"approved_time"`
	SourceIp                 string    `json:"source_ip"`
	SourceBrowser            string    `json:"source_browser"`
	RejectedDate             time.Time `json:"rejected_date" gorm:"type:date;null"`
	RejectedTime             time.Time `json:"rejected_time"`
	RejectedbyID             uuid.UUID `json:"rejectedby_id"`

	FkCreatedbyID      User `gorm:"foreignKey:CreatedbyID;"`
	FkApprovedbyID     User `gorm:"foreignKey:ApprovedbyID;"`
	FkRejectedbyID     User `gorm:"foreignKey:RejectedbyID;"`
	FkLastmodifiedbyID User `gorm:"foreignKey:LastmodifiedbyID;"`
}
