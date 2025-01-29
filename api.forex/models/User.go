package models

import "github.com/google/uuid"

type Fx_User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}
