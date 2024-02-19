package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Credential struct {
	ID        string         `json:"id" gorm:"primaryKey;column:id"`
	AccountID string         `json:"account_id" gorm:"column:account_id;not null"`
	Account   *Account       `json:"-" gorm:"foreignKey:AccountID"`
	Platform  Platform       `json:"platform" gorm:"column:platform;not null"`
	Type      PlatformType   `json:"type" gorm:"column:type;not null"`
	IsActive  int            `json:"is_active" gorm:"column:is_active;default:1"`
	Meta      Map            `json:"meta" gorm:"column:meta;type:jsonb;default:'{}'"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;type:timestamptz"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;type:timestamptz"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:timestamptz"`
}

func (Credential) TableName() string {
	return "credentials"
}

func (c *Credential) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now().UTC()
	}
	if c.UpdatedAt.IsZero() {
		c.UpdatedAt = time.Now().UTC()
	}
	return nil
}

type Platform string
type PlatformType string

const (
	// email platforms
	MAILJET Platform = "mailjet"
	MAILGUN Platform = "mailgun"
)

const (
	EMAIL PlatformType = "email"
	SMS   PlatformType = "sms"
	PUSH  PlatformType = "push"
)
