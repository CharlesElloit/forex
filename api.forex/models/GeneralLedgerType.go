package models

import (
	"time"

	"github.com/google/uuid"
)

type Fx_GL_Type struct {
	ID               uuid.UUID  `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"` //SERIAL NOT NULL PRIMARY KEY UNIQUE
	Name             string     `json:"name"`                                                            //VARCHAR(100) NOT NULL
	ShortName        string     `json:"short_name"`                                                      //CHAR(2) NOT NULL
	Description      string     `json:"description"`                                                     //VARCHAR(200) NULL
	CreatedDate      time.Time  `json:"created_date"`                                                    //DATE NOT NULL DEFAULT CURRENT_DATE
	CreatedByID      string     `json:"createdby_id"`                                                    //INTEGER REFERENCES fx_user(id) NOT NULL
	LastmodifiedDate *time.Time `json:"lastmodified_date"`                                               //DATE NULL
	LastmodifiedByID *string    `json:"lastmodifiedby_id"`                                               //INTEGER REFERENCES fx_user(id)
	IsApproved       *bool      `json:"is_approved"`                                                     //BOOLEAN NOT NULL DEFAULT FALSE
	ApprovedByID     *string    `json:"approvedby_id"`                                                   //INTEGER REFERENCES fx_user(id) NULL
}
