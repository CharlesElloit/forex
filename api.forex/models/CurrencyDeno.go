package models

import (
	"time"

	"github.com/google/uuid"
)

type Fx_Currency_Deno struct {
	ID                    uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`                     //SERIAL PRIMARY KEY NOT NULL UNIQUE
	DenoName              string    `json:"deno_name" gorm:"type:varchar(100);not null"`                                         // VARCHAR(100) NOT NULL
	DenoAmount            float32   `json:"deno_amount" gorm:"not null"`                                                         // FLOAT NOT NULL
	IsDenoActive          bool      `json:"is_deno_active" gorm:"default:false"`                                                 //BOOLEAN NOT NULL DEFAULT FALSE
	DenoActivitionDate    string    `json:"deno_activition_date"`                                                                // DATE NULL
	DenoActivitionTime    time.Time `json:"deno_activition_time"`                                                                //VARCHAR(8) NULL
	DenoActivitedByID     uuid.UUID `json:"deno_activitedby_id"`                                                                 //INTEGER REFERENCES fx_user(id) NULL
	DenoInactiveDate      string    `json:"deno_inactive_date"`                                                                  // DATE NULL
	DenoInactiveTime      time.Time `json:"deno_inactivated_time"`                                                               //VARCHAR(8) NULL
	DenoInactiveByID      uuid.UUID `json:"deno_inactivateby_id"`                                                                //INTEGER REFERENCES fx_user(id) NULL
	DenoFigure            int       `json:"deno_figure" gorm:"not null"`                                                         //INTEGER NOT NULL
	CurrencyID            uuid.UUID `json:"currency_id" gorm:"type:uuid;not null"`                                               // INTEGER REFERENCES fx_currency(id) UNIQUE NOT NULL
	AllowDecimalPrecision bool      `json:"allow_decimal_precision" gorm:"default:false"`                                        //BOOLEAN NOT NULL DEFAULT FALSE
	DecimalPrecisionDigit int       `json:"decimal_precision_digit" gorm:"type:int"`                                             //INTEGER NULL
	IsCoin                bool      `json:"is_coin" gorm:"default:false"`                                                        // BOOLEAN NOT NULL DEFAULT FALSE
	CreatedDate           string    `json:"created_date" gorm:"type:date;not null;default:current_date"`                         // DATE NOT NULL DEFAULT CURRENT_DATE
	CreatedTime           time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"` //VARCHAR(8) NOT NULL DEFAULT TO_CHAR(NOW(), 'HH24:MI:SS AM')
	CreatedByID           uuid.UUID `json:"createdby_id" gorm:"not null"`                                                        // INTEGER REFERENCES fx_user(id) NOT NULL
	LastmodifiedDate      string    `json:"lastmodified_date"`                                                                   // DATE NULL
	LastmodifiedTime      time.Time `json:"lastmodified_time"`                                                                   //VARCHAR(8) NULL
	LastmodifiedByID      uuid.UUID `json:"lastmodifiedby_id"`                                                                   // INTEGER REFERENCES fx_user(id) NULL
	IsDeleted             bool      `json:"is_deleted"`                                                                          // BOOLEAN NOT NULL DEFAULT FALSE
	DeletedDate           string    `json:"deleted_date"`                                                                        // DATE NULL
	DeletedTime           time.Time `json:"deleted_time"`                                                                        //VARCHAR(8) NULL
	DeletedByID           uuid.UUID `json:"deletedby_id"`                                                                        //INTEGER REFERENCES fx_user(id) NULL
	IsActive              bool      `json:"is_active"`                                                                           // BOOLEAN NULL
	IsApproved            bool      `json:"is_approved"`                                                                         // BOOLEAN NOT NULL DEFAULT FALSE
	ApprovedTime          time.Time `json:"approved_time"`                                                                       //VARCHAR(8) NULL
	ApprovedByID          uuid.UUID `json:"approvedby_id"`                                                                       // INTEGER REFERENCES fx_user(id) NULL
	SourceIp              string    `json:"source_ip" gorm:"type:varchar(15);not null"`                                          //VARCHAR(15) NOT NULL
	SourceBrowser         string    `json:"source_browser" gorm:"type:varchar(30);not null"`                                     //VARCHAR(15) NOT NULL
	RejectedDate          string    `json:"rejected_date"`                                                                       //DATE NULL
	RejectedTime          string    `json:"rejected_time"`                                                                       //VARCHAR(8) NULL
	RejectedByID          uuid.UUID `json:"rejectedby_id"`                                                                       //INTEGER REFERENCES fx_user(id) NULL
}
