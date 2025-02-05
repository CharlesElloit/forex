package models

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	ResourceName     string    `json:"resource_name" gorm:"type:varchar(100);not null"`
	ModuleID         uuid.UUID `json:"module_id" gorm:"type:uuid"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime      time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `gorm:"not null;index"`
	LastmodifiedDate string    `json:"lastmodified_date" gorm:"type:date;null"`
	LastmodifiedTime time.Time `json:"lastmodified_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	LastmodifiedbyID uuid.UUID `gorm:"not null;index"`

	FkModuleID         Module `gorm:"foreignKey:ModuleID;"`
	FkCreatedbyID      User   `gorm:"foreignKey:CreatedbyID;"`
	FkLastmodifiedbyID User   `gorm:"foreignKey:LastmodifiedbyID;"`
}
