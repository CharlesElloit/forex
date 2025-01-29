package models

import (
	"time"

	"github.com/google/uuid"
)

type Fx_GL_Total struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`            //SERIAL NOT NULL PRIMARY KEY UNIQUE
	BranchID         string    `json:"branch_id"`                                                                  //INTEGER REFERENCES fx_branch(id) NOT NULL
	HostBranch       string    `json:"host_branch"`                                                                //INTEGER REFERENCES fx_branch(id) NOT NULL
	GlName           string    `json:"gl_name"`                                                                    //VARCHAR(100) NOT NULL
	AccountCode      string    `json:"account_code"`                                                               //NUMERIC(10) NOT NULL
	AccountType      string    `json:"account_type"`                                                               //CHAR(2) NOT NULL
	BalanceType      string    `json:"balance_type"`                                                               //credit_debit_enum_type NOT NULL --This describes the type of balance (debit or credit)
	GlGroupID        string    `json:"gl_group_id"`                                                                //INTEGER REFERENCES fx_gl_group(id) NOT NULL
	CurrencyID       string    `json:"currency_id"`                                                                //INTEGER REFERENCES fx_currency(id) NOT NULL
	TransactionDate  string    `json:"transaction_date"`                                                           //DATE NOT NULL DEFAULT CURRENT_DATE
	TransactionTime  time.Time `json:"transaction_time"`                                                           // VARCHAR(8) NOT NULL DEFAULT TO_CHAR(NOW(), 'HH24:MI:SS')
	OpeningBalance   float32   `json:"opening_balance" gorm:"default:0"`                                           // NUMERIC(19, 2) NOT NULL DEFAULT 0
	ClosingBalance   float32   `json:"closing_balance" gorm:"default:0"`                                           // NUMERIC(19, 2) NOT NULL DEFAULT 0
	Rate             float32   `json:"rate"`                                                                       // FLOAT NOT NULL DEFAULT 0
	LcEquivalent     float32   `json:"lc_equivalent"`                                                              // NUMERIC(19, 2) NULL
	DebitMovement    float32   `json:"debit_movement"`                                                             // NUMERIC(19, 2) NOT NULL DEFAULT 0
	CreditMovement   float32   `json:"credit_movement"`                                                            // NUMERIC(19, 2) NOT NULL DEFAULT 0
	CreatedDate      string    `json:"created_date" gorm:"default:current_date"`                                   //DATE NOT NULL DEFAULT CURRENT_DATE
	CreatedTime      time.Time `json:"created_time" gorm:"type:varchar(8);default:to_char(now(),'HH24:MI:SS AM')"` // VARCHAR(8) NOT NULL DEFAULT TO_CHAR(NOW(), 'HH24:MI:SS')
	CreatedByID      string    `json:"createdby_id"`                                                               // INTEGER REFERENCES fx_user(id) NOT NULL
	LastmodifiedDate string    `json:"lastmodified_date"`                                                          // DATE NULL
	LastmodifiedTime time.Time `json:"lastmodified_time"`                                                          // VARCHAR(8) NULL
	LastmodifiedByID string    `json:"lastmodifiedby_id"`                                                          // INTEGER REFERENCES fx_user(id) NULL
}
