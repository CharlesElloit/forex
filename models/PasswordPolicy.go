package models

import (
	"time"

	"github.com/google/uuid"
)

type PasswordPolicy struct {
	ID                            uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	PolicyName                    string    `json:"policy_name" gorm:"type:varchar(100);not null;unique"`
	Description                   string    `json:"description" gorm:"type:varchar(200)"`
	MinPasswordLength             int       `json:"min_password_length" gorm:"type:int;default:0"`
	MaxPasswordLength             int       `json:"max_password_length" gorm:"type:int;default:0"`
	PasswordExpiryDays            int       `json:"password_expiry_days" gorm:"type:int;default:0"`
	GracePeriodDays               int       `json:"grace_period_days" gorm:"type:int;default:0"`
	EnforcePasswordHistorylimit   bool      `json:"enforce_password_history_limit" gorm:"type:boolean;default:false"`
	PasswordHistoryLimit          int       `json:"password_history_limit" gorm:"type:int;default:0"`
	EnableComplexity              bool      `json:"enable_complexity" gorm:"type:boolean;default:false"`
	AllowRepeatedCharacters       bool      `json:"allow_repeated_characters" gorm:"type:boolean;default:true"`
	MinUppercase                  int       `json:"min_uppercase" gorm:"type:int;default:0"`
	MinLowercase                  int       `json:"min_lowercase" gorm:"type:int;default:0"`
	MinNumbers                    int       `json:"min_numbers" gorm:"type:int;default:0"`
	MinSpecialCharacters          int       `json:"min_special_characters" gorm:"type:int;default:0"`
	EnforceDictionaryCheck        bool      `json:"enforce_dictionary_check" gorm:"type:boolean;default:false"`
	DictionarySourceFilePath      string    `json:"dictionary_source_file_path" gorm:"type:varchar(255);default:null"`
	DictionarySourceApi           string    `json:"dictionary_source_api" gorm:"type:varchar(255);default:null"`
	LockoutDurationMinutes        int       `json:"lockout_duration_minutes" gorm:"type:int;default:0"`
	LockoutThreshold              int       `json:"lockout_threshold" gorm:"type:int;default:0"`
	RequireMfaAfterFailedAttempts int       `json:"require_mfa_after_failed_attempts" gorm:"type:int;default:0"`
	IsPolicyActive                bool      `json:"is_policy_active" gorm:"type:boolean;default:false"`
	EnablePermanentLockout        bool      `json:"enable_permanent_lockout" gorm:"type:boolean;default:false"`
	MinLockoutCount               int       `json:"min_lockout_count" gorm:"type:int;default:0"`
	CreatedDate                   time.Time `json:"created_date" gorm:"type:varchar(15);not null;default:current_date"`
	CreatedTime                   time.Time `json:"created_time" gorm:"type:varchar(8);not null;default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID                   uuid.UUID `json:"createdby_id" gorm:"type:uuid;foreignKey:UserID"`
	LastmodifiedDate              string    `json:"lastmodified_date"`
	LastmodifiedTime              time.Time `json:"lastmodified_time"`
	LastmodifiedbyID              uuid.UUID `json:"lastmodifiedby_id" gorm:"type:uuid;foreignKey:UserID;default:null"`
	SourceIp                      string    `json:"source_ip" gorm:"type:varchar(15);not null"`
	SourceBrowser                 string    `json:"source_browser" gorm:"type:varchar(30);not null"`
	UserID                        uuid.UUID
}
