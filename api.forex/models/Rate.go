package models

import (
	"time"

	"github.com/google/uuid"
)

type Fx_Rate struct {
	ID               uuid.UUID  `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	SellingRate      float32    `json:"selling_rate"`
	BuyingRate       float32    `json:"buying_rate"`
	AverageRate      float32    `json:"average_rate"`
	MidRate          float32    `json:"mid_rate"`
	CreatedDate      string     `json:"created_date" gorm:"default:current_date"`
	CreatedTime      string     `json:"created_time" gorm:"default:to_char(now(), 'HH24:MI:SS AM')"`
	CreatedByID      string     `json:"createdby_id"` // reference to user
	LastModifiedDate string     `json:"lastmodified_date"`
	LastModifiedByID string     `json:"lastmodifiedby_id"` // reference to user
	IsApproved       bool       `json:"is_approved"`
	ApprovedByID     string     `json:"approvedby_id"` // reference to user
	ApprovedTime     string     `json:"approved_time"`
	RejectedDate     *string    `json:"rejected_date"` //DATE NULL
	RejectedTime     *time.Time `json:"rejected_time"` //VARCHAR(8) NULL
	RejectedByID     *string    `json:"rejectedby_id"` //INTEGER REFERENCES fx_user(id) NULL
}
