package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Key struct {
	ID        string `json:"id" gorm:"primaryKey;column:id"`
	AccountID string `json:"account_id" gorm:"column:account_id;not null"`
	Name      string `json:"name" gorm:"column:name;not null"`

	Public    string    `json:"public" gorm:"column:public;not null"` // only this in use
	Private   string    `json:"private" gorm:"column:private;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamptz"`

	Account *Account `json:"-" gorm:"foreignKey:AccountID"`
}

func (Key) TableName() string {
	return "keys"
}

func (k *Key) BeforeCreate(tx *gorm.DB) error {
	if k.ID == "" {
		k.ID = uuid.New().String()
	}
	if k.CreatedAt.IsZero() {
		k.CreatedAt = time.Now().UTC()
	}
	return nil
}

// I'm not comfortable with this method of generating keys
func (k *Key) GenerateKey() error {
	k.Public = "public-" + time.Now().Format("20060102150405MonMSTJan") + string(uuid.New().NodeID())
	k.Private = "private-" + time.Now().Format("20060102150405MSTJan") + string(uuid.New().NodeID())
	return nil
}
