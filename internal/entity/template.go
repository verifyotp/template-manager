package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Template struct {
	ID        string `json:"id" gorm:"primaryKey;column:id"`
	AccountID string `json:"account_id" gorm:"column:account_id;not null"`

	Name        string `json:"name" gorm:"column:name;not null"`
	Slug        string `json:"slug" gorm:"column:slug;not null"` // many templates can have the same slug but different versions
	Version     uint64 `json:"version" gorm:"column:version;not null;default:1"`
	Location    string `json:"location" gorm:"column:location;not null"` // location of the template [url link]
	ContentType string `json:"content_type" gorm:"column:content_type;not null"`
	Vars        Map    `json:"vars" gorm:"column:vars;type:jsonb;not null"` // pre-existing values are treated as default values

	Active bool `json:"active" gorm:"column:active;not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;type:timestamptz"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;type:timestamptz"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:timestamptz"`

	Account *Account `json:"-" gorm:"foreignKey:AccountID"`
}

func (Template) TableName() string {
	return "templates"
}

func (t *Template) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now().UTC()
	}
	if t.UpdatedAt.IsZero() {
		t.UpdatedAt = time.Now().UTC()
	}
	return nil
}

type TemplateSync struct {
	ID       string `json:"id" gorm:"primaryKey;column:id"`
	Provider string `json:"provider" gorm:"column:provider;not null"`

	Identifier  string `json:"identifier" gorm:"column:identifier;not null"` // from the provider  [ e.g the id of the template from the provider]
	Version     string `json:"version" gorm:"column:version;not null;default:'1'"`
	Location    string `json:"location" gorm:"column:location;not null"` // location of the template [url link]
	ContentType string `json:"content_type" gorm:"column:content_type;not null"`
	Vars        Map    `json:"vars" gorm:"column:vars;type:jsonb;not null"` // pre-existing values are treated as default values

	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;type:timestamptz"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;type:timestamptz"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:timestamptz"`
}

func (TemplateSync) TableName() string {
	return "template_syncs"
}
