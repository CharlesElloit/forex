package models

import (
	"time"

	"github.com/google/uuid"
)

type GL struct {
	ID                   uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	BranchID             uuid.UUID `json:"branch_id"`
	GlTypeID             uuid.UUID `json:"gl_type_id" gorm:"type:uuid;not null"`
	GlGroupID            uuid.UUID `json:"gl_group_id" gorm:"type:uuid;not null"`
	GlCode               int       `json:"gl_code" gorm:"type:int;not null"`
	GlFullCode           int       `json:"gl_full_code"`
	GlName               string    `json:"gl_name" gorm:"type:varchar(100)"`
	Description          string    `json:"description"`
	CurrencyID           string    `json:"currency_id" gorm:"type:uuid;not null"`
	DebitCredit          string    `json:"debit_credit"`
	OpeningBalance       float32   `json:"opening_balance" gorm:"type:float;not null;default:0.0"`
	ClosingBalance       float32   `json:"closing_balance" gorm:"type:float;not null;default:0.0"`
	OpenedDate           time.Time `json:"opened_date" gorm:"type:date;not null;default:current_date"`
	ClosedDate           string    `json:"closed_date" gorm:"type:date;null"`
	ManualEntriesAllowed bool      `json:"manual_entries_allowed" gorm:"type:boolean;not null;default:false"`
	AllowReconcilation   bool      `json:"allow_reconcilation" gorm:"type:boolean;not null;default:false"`
	AllowRevalution      bool      `json:"allow_revalution" gorm:"type:boolean;not null;default:false"`
	CreatedDate          time.Time `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime          string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID          uuid.UUID `json:"createdby_id" gorm:"type:uuid;not null"`
	LastmodifiedDate     string    `json:"lastmodified_date" gorm:"type:date;null"`
	LastmodifiedTime     time.Time `json:"lastmodified_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	LastmodifiedbyID     uuid.UUID `json:"lastmodifiedby_id" gorm:"type:uuid;null"`
	IsApproved           bool      `json:"is_approved" gorm:"type:boolean;null"`
	ApprovedbyID         string    `json:"approvedby_id" gorm:"type:uuid;null"`
	ApprovedTime         time.Time `json:"approved_time"`
	RejectedDate         string    `json:"rejected_date"`
	RejectedTime         string    `json:"rejected_time"`
	RejectedbyID         uuid.UUID `json:"rejectedby_id"`

	FkGlTypeID         GlType  `gorm:"foreignKey:GlTypeID"`
	FkBranchID         Branch  `gorm:"foreignKey:BranchID"`
	FkGlGroupCodeID    GlGroup `gorm:"foreignKey:GlGroupID"`
	FkCreatedbyID      User    `gorm:"foreignKey:CreatedbyID;"`
	FkApprovedbyID     User    `gorm:"foreignKey:ApprovedbyID;"`
	FkRejectedbyID     User    `gorm:"foreignKey:RejectedbyID;"`
	FkLastmodifiedbyID User    `gorm:"foreignKey:LastmodifiedbyID;"`
}
