package models

import (
	"time"

	"github.com/google/uuid"
)

type Fx_Branch struct {
	ID               uuid.UUID  `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`            //SERIAL PRIMARY KEY NOT NULL
	Name             string     `json:"name" gorm:"type:varchar(30)"`                                               //VARCHAR(30) NOT NULL
	Code             int        `json:"code"`                                                                       //INTEGER NOT NULL
	OpenDate         string     `json:"open_date" gorm:"default:current_date"`                                      //DATE NOT NULL DEFAULT CURRENT_DATE
	CloseDate        *string    `json:"close_date"`                                                                 //DATE NULL
	Address          string     `json:"address"`                                                                    //VARCHAR(200) NOT NULL
	StartTime        time.Time  `json:"start_time"`                                                                 //VARCHAR(8) NOT NULL
	EndTime          time.Time  `json:"end_time"`                                                                   //VARCHAR(8) NOT NULL
	Type             string     `json:"type"`                                                                       //CHAR(1) NOT NULL --regular or head office
	Contact          string     `json:"contact" gorm:"type:varchar(14)"`                                            //VARCHAR(14) NULL
	EmailAddres      string     `json:"email_address"`                                                              //VARCAHR(255) NULL
	CreatedDate      string     `json:"created_date" gorm:"default:current_date"`                                   //DATE NOT NULL DEFAULT CURRENT_DATE
	CreatedTime      time.Time  `json:"created_time" gorm:"type:varchar(8);default:to_char(now(),'HH24:MI:SS AM')"` //VARCHAR(8) NOT NULL DEFAULT TO_CHAR(NOW() 'HH24:MI:SS AM')
	CreatedByID      string     `json:"createdby_id"`                                                               //INTEGER REFERENCES fx_user(id) NOT NULL
	LastmodifiedDate *string    `json:"lastmodified_date"`                                                          //DATE NULL
	LastmodifiedTime *time.Time `json:"lastmodified_time"`                                                          //VARCHAR(8) NULL
	LastmodifiedByID string     `json:"lastmodifiedby_id"`                                                          //INTEGER REFERENCES fx_user(id) NULL
	ApprovedDate     string     `json:"approved_date"`                                                              //DATE NULL
	ApprovedTime     time.Time  `json:"approved_time"`                                                              //VARCHAR(8) NULL
	IsApproved       bool       `json:"is_approved"`                                                                //BOOLEAN NOT NULL DEFAULT false
	ApprovedByID     string     `json:"approvedby_id"`                                                              //INTEGER REFERENCES fx_user(id) NULL
	SourceIp         string     `json:"source_ip"`                                                                  //VARCHAR(15) NOT NULL
	SourceBrowser    string     `json:"source_browser"`                                                             //VARCHAR(15) NOT NULL
	RejectedDate     string     `json:"rejected_date"`                                                              //DATE NULL
	RejectedTime     string     `json:"rejected_time"`                                                              //VARCHAR(8) NULL
	RejectedByID     string     `json:"rejectedby_id"`                                                              //INTEGER REFERENCES fx_user(id) NULL
}
