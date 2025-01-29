package models

import (
	"time"

	"github.com/google/uuid"
)

type Fx_Charges_Type struct {
	ID               uuid.UUID  `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"` //SERIAL PRIMARY KEY NOT NULL UNIQUE
	Name             string     `json:"name"`                                                            //VARCHAR(100) NOT NULL
	Description      string     `json:"description"`                                                     //VARCHAR(200) NULL
	CreatedDate      string     `json:"created_date" gorm:"default:current_date"`                        //DATE NOT NULL DEFAULT CURRENT_DATE
	CreatedTime      string     `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`      //VARCHAR(8) NOT NULL DEFAULT TO_CHAR(NOW() 'HH24:MI:SS AM')
	CreatedByID      string     `json:"createdby_id"`                                                    //INTEGER REFERENCES fx_user(id) NOT NULL
	LastmodifiedDate string     `json:"lastmodified_date"`                                               //DATE NULL
	LastmodifiedTime time.Time  `json:"lastmodified_time"`                                               //VARCHAR(8) NULL
	LastmodifiedByID string     `json:"lastmodifiedby_id"`                                               //INTEGER REFERENCES fx_user(id) NULL
	IsActive         bool       `json:"is_active"`                                                       //BOOLEAN NOT NULL DEFAULT FALSE
	IsApproved       bool       `json:"is_approved"`                                                     //BOOLEAN NOT NULL DEFAULT FALSE
	ApprovedTime     *time.Time `json:"approved_time"`                                                   //VARCHAR(8) NULL
	ApprovedByID     *string    `json:"approvedby_id"`                                                   //INTEGER REFERENCES fx_user(id) NULL
	SourceIp         string     `json:"source_ip"`                                                       //VARCHAR(15) NOT NULL
	SourceBrowser    string     `json:"source_browser"`                                                  //VARCHAR(15) NOT NULL
	RejectedDate     *string    `json:"rejected_date"`                                                   //DATE NULL
	RejectedTime     *time.Time `json:"rejected_time"`                                                   //VARCHAR(8) NULL
	RejectedByID     *string    `json:"rejectedby_id"`                                                   //INTEGER REFERENCES fx_user(id) NULL
}
