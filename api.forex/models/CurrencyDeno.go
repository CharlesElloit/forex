package models

import (
	"time"

	"github.com/google/uuid"
)

type CurrencyDeno struct {
	ID                    uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	DenoName              string    `json:"deno_name" gorm:"type:varchar(100);not null"`
	DenoAmount            float32   `json:"deno_amount" gorm:"not null"`
	IsDenoActive          bool      `json:"is_deno_active" gorm:"default:false"`
	DenoActivitionDate    string    `json:"deno_activition_date"`
	DenoActivitionTime    time.Time `json:"deno_activition_time"`
	DenoActivitedbyID     uuid.UUID `json:"deno_activitedby_id"`
	DenoInactiveDate      string    `json:"deno_inactive_date"`
	DenoInactiveTime      time.Time `json:"deno_inactivated_time"`
	DenoInactivebyID      uuid.UUID `json:"deno_inactivateby_id"`
	DenoFigure            int       `json:"deno_figure" gorm:"not null"`
	CurrencyID            uuid.UUID `json:"currency_id" gorm:"type:uuid;not null"`
	AllowDecimalPrecision bool      `json:"allow_decimal_precision" gorm:"default:false"`
	DecimalPrecisionDigit int       `json:"decimal_precision_digit" gorm:"type:int"`
	IsCoin                bool      `json:"is_coin" gorm:"default:false"`
	CreatedDate           time.Time `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime           time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID           uuid.UUID `gorm:"not null;index"`
	LastmodifiedDate      string    `json:"lastmodified_date" gorm:"type:date;null"`
	LastmodifiedTime      time.Time `json:"lastmodified_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	LastmodifiedbyID      uuid.UUID `gorm:"not null;index"`
	IsDeleted             bool      `json:"is_deleted"`
	DeletedDate           string    `json:"deleted_date"`
	DeletedTime           time.Time `json:"deleted_time"`
	DeletedbyID           uuid.UUID `json:"deletedby_id"`
	IsActive              bool      `json:"is_active"`
	IsApproved            bool      `json:"is_approved"`
	ApprovedTime          time.Time `json:"approved_time"`
	ApprovedbyID          uuid.UUID `json:"approvedby_id"`
	SourceIp              string    `json:"source_ip" gorm:"type:varchar(15);not null"`
	SourceBrowser         string    `json:"source_browser" gorm:"type:varchar(30);not null"`
	RejectedDate          string    `json:"rejected_date"`
	RejectedTime          string    `json:"rejected_time"`
	RejectedbyID          uuid.UUID `json:"rejectedby_id"`

	FkCurrencyID        Currency `gorm:"foreignKey:CurrencyID;"`
	FkCreatedbyID       User     `gorm:"foreignKey:CreatedbyID;"`
	FkApprovedbyID      User     `gorm:"foreignKey:ApprovedbyID;"`
	FkDeletedbyID       User     `gorm:"foreignKey:DeletedbyID;"`
	FkRejectedbyID      User     `gorm:"foreignKey:RejectedbyID;"`
	FkDenoInactivebyID  User     `gorm:"foreignKey:DenoInactivebyID;null"`
	FkDenoActivitedbyID User     `gorm:"foreignKey:DenoActivitedbyID;null"`
	FkLastmodifiedbyID  User     `gorm:"foreignKey:LastmodifiedbyID;"`
}
