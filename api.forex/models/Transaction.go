package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                uuid.UUID  `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	Srno              int        `json:"srno"`
	BranchId          int        `json:"branch_id"`
	AccountCode       string     `json:"account_code"`
	Narration         string     `json:"narration" gorm:"type:varchar(200)"`
	AccountType       string     `json:"account_type"`
	DebitCredit       string     `json:"debit_credit"`
	Amount            float32    `json:"amount"`
	LcEqvantentAmount float32    `json:"lc_eqvantent_amount"`
	Rate              float32    `json:"rate"`
	VoucherNo         int        `json:"voucher_no"`
	CurrencyID        string     `json:"currency_id"`
	TransactionType   string     `json:"transaction_type"`
	TransactionMode   string     `json:"transaction_mode"`
	TransactionDate   time.Time  `json:"transaction_date" gorm:"default:current_date"`
	TransactionTime   time.Time  `json:"transaction_time" gorm:"type:varchar(8);default:to_char(now(),'HH24:MI:SS AM')"`
	IsDeleted         bool       `json:"is_deleted" gorm:"default:false"`
	IsApproved        bool       `json:"is_approved"`
	CreatedDate       string     `json:"created_date" gorm:"default:current_date"`
	CreatedByID       string     `josn:"createdby_id"`
	LastModifiedDate  *string    `json:"lastmodified_date"`
	LastModifiedByID  *string    `json:"lastmodifiedby_id"`
	ApprovedByID      *string    `json:"approvedby_id"`
	ApprovedTime      *time.Time `json:"approved_time"`
	SourceIP          string     `json:"source_ip"`
	SourceBrowser     string     `json:"source_browser"`
	RejectedDate      *string    `json:"rejected_date"`
	RejectedTime      *time.Time `json:"rejected_time"`
	RejectedByID      *string    `json:"rejectedby_id"`
}
