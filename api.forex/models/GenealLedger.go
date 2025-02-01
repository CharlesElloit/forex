package models

import (
	"time"

	"github.com/google/uuid"
)

type GL struct {
	ID                   uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	BranchID             string    `json:"branch_id"`        // reference to the fx_branch table(id)
	GLTypeID             string    `json:"gl_type_id"`       // reference to the fx_gl_type(id)
	GLGroupCodeID        string    `json:"gl_group_code_id"` //reference to the fx_gl_group(id) table
	GLCode               int       `json:"gl_code"`
	GLFullCode           int       `json:"gl_full_code"`
	GLName               string    `json:"gl_name"`
	Description          string    `json:"description"`
	CurrencyID           string    `json:"currency_id"`  // reference to the fx_currency(id)
	DebitCredit          string    `json:"debit_credit"` // enum type credit_debit_enum_type
	OpeningBalance       float32   `json:"opening_balance"`
	ClosingBalance       float32   `json:"closing_balance"`
	OpenedDate           string    `json:"opened_date"`
	ClosedDate           string    `json:"closed_date"`
	ManualEntriesAllowed bool      `json:"manual_entries_allowed"`
	AllowReconcilation   bool      `json:"allow_reconcilation"`
	AllowRevalution      bool      `json:"allow_revalution"`
	CreatedDate          string    `json:"created_date"`
	CreatedTime          string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedByID          string    `sjon:"createdby_id"` // reference to the fx_user(id)
	LastmodifiedDate     string    `json:"lastmodified_date"`
	LastmodifiedByID     string    `json:"lastmodifiedby_id"` // reference to the fx_user(id)
	IsApproved           bool      `json:"is_approved"`
	ApprovedByID         string    `json:"approvedby_id"` // reference to the fx_user(id)
	ApprovedTime         time.Time `json:"approved_time"`
}
