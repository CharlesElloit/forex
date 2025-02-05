package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	UserID           User      `json:"user_id" gorm:"type:uuid;not null"`
	RoleID           uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime      time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `gorm:"not null;index"`
	LastmodifiedDate string    `json:"lastmodified_date" gorm:"type:date;null"`
	LastmodifiedTime time.Time `json:"lastmodified_time" gorm:"type:varchar(8);null"`
	LastmodifiedbyID uuid.UUID `gorm:"null;index"`

	FkCreatedbyID      User `gorm:"foreignKey:CreatedbyID;"`
	FkLastmodifiedbyID User `gorm:"foreignKey:LastmodifiedbyID;"`
	FkUserID           User `gorm:"foreignKey:UserID;"`
	FkRoleID           Role `gorm:"foreignKey:RoleID;"`
}
