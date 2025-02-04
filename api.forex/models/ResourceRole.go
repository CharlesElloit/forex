package models

import (
	"time"

	"github.com/google/uuid"
)

type ResourceRole struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	ResourceID       uuid.UUID `josn:"resource_id" gorm:"type:uuid;not null"`
	RoleID           uuid.UUID `josn:"role_id" gorm:"type:uuid;not null"`
	CanRead          bool      `josn:"can_read" gorm:"type:boolean;not null;default:false"`
	CanAdd           bool      `josn:"can_add"  gorm:"type:boolean;not null;default:false"`
	CanEdit          bool      `josn:"can_edit" gorm:"type:boolean;not null;default:false"`
	CanCreate        bool      `josn:"can_create" gorm:"type:boolean;not null;default:false"`
	CanDelete        bool      `josn:"can_delete" gorm:"type:boolean;not null;default:false"`
	CreatedTime      string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `sjon:"createdby_id"`
	LastmodifiedDate string    `json:"lastmodified_date"`
	LastmodifiedTime time.Time `json:"lastmodified_time" gorm:"type:varchar(8);default:to_char(now(), 'HH24:MI:SS AM')"`
	LastmodifiedbyID uuid.UUID `json:"lastmodifiedby_id"`

	FkRoleID           Role     `gorm:"foreignKey:RoleID"`
	FkResourceID       Resource `gorm:"foreignKey:ResourceID"`
	FkCreatedbyID      User     `gorm:"foreignKey:CreatedbyID;"`
	FkLastmodifiedbyID User     `gorm:"foreignKey:LastmodifiedbyID;"`
}
