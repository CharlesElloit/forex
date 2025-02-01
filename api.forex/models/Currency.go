package models

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	ID                       uuid.UUID  `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	CurrencyName             string     `json:"currency_name" gorm:"unique"`
	CurrencyCode             string     `json:"currency_code" gorm:"unique"`
	DecimalPlaces            int        `json:"decimal_places"`
	DisplaySymbol            string     `json:"display_symbol"`
	CurrencyStatus           string     `json:"currency_status"`
	InternationlizedNameCode string     `json:"internationalized_name_code"`
	ActivitionDate           string     `json:"activition_date"`
	InactiveDate             string     `json:"inactive_date"`
	CreatedDate              *time.Time `json:"created_date" gorm:"default:current_date"`
	CreatedTime              string     `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedByID              string     `json:"createdby_id"` // reference to user table
	LastModifiedDate         *time.Time `json:"lastmodified_date"`
	LastModifiedByID         string     `json:"lastmodifiedby_id"` // reference to user table
	IsApproved               bool       `json:"is_approved" gorm:"default:false"`
	ApprovedByID             string     `json:"approvedby_id"` // reference to user table
	ApprovedTime             *time.Time `json:"approved_time"`
}
