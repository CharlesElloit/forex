package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	RoleName         string    `json:"role_name" gorm:"type:varchar(200);not null;unique"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:varchar(15);not null;default:current_date"`
	CreatedTime      time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `gorm:"not null;index"`
	LastmodifiedDate string    `json:"lastmodified_date"`
	LastmodifiedTime time.Time `json:"lastmodified_time"`
	LastmodifiedbyID uuid.UUID `gorm:"not null;index"`

	FkCreatedbyID      User `gorm:"foreignKey:CreatedbyID;"`
	FkLastmodifiedbyID User `gorm:"foreignKey:LastmodifiedbyID;"`
}
