package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID                      uuid.UUID      `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	FirstName               string         `json:"firstName"`
	LastName                string         `json:"lastName"`
	Email                   string         `json:"email" gorm:"type:varchar(255);unique" validator:"email;required"`
	Password                string         `json:"password"`
	PasswordExpiryDate      time.Time      `json:"password_expiry_date" gorm:"type:date"`
	PasswordLastUpdatedDate time.Time      `json:"password_last_updated_date" gorm:"type:date"`
	AccountLocked           bool           `json:"account_locked" gorm:"type:boolean;not null;default:false"`
	AccountActive           bool           `json:"account_active" gorm:"type:boolean;not null;default:true"`
	Roles                   []*Role        `json:"roles" gorm:"many2many:user_roles"`
	UserScopes              datatypes.JSON `json:"user_scopes"`
	CreatedDate             time.Time      `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime             time.Time      `json:"created_time" gorm:"type:varchar(12);not null;default:to_char(now(), 'HH24:MI:SS AM')"`
	CreatedbyID             uuid.UUID      `gorm:"type:uuid;not null;index"`
	LastmodifiedDate        time.Time      `json:"lastmodified_date"`
	LastmodifiedTime        string         `json:"lastmodified_time" gorm:"type:varchar(12);null"`
	LastmodifiedbyID        uuid.UUID      `json:"lastmodifiedby_id"`
	ApprovedByID            uuid.UUID      `json:"approvedby_id"`
	ApprovedTime            string         `json:"approved_time" gorm:"type:varchar(12);null"`
	SourceIP                string         `json:"source_ip" gorm:"type:varchar(15);not null"`
	SourceBrowser           string         `json:"source_browser" gorm:"type:varchar(255);not null"`
	RejectedDate            time.Time      `json:"rejected_date"`
	RejectedTime            time.Time      `json:"rejected_time"`
	RejectedbyID            uuid.UUID      `json:"rejectedby_id" gorm:"type:uuid;null"`

	FkCreatedbyID      *User `gorm:"foreignKey:CreatedbyID"`
	FkRejectedbyID     *User `gorm:"foreignKey:RejectedbyID;"`
	FkLastmodifiedbyID *User `gorm:"foreignKey:LastmodifiedbyID;"`
}
