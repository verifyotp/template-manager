package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	bycrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	ID             string     `json:"id" gorm:"primaryKey;column:id"`
	Email          string     `json:"email" gorm:"unique;column:email"`
	HashSalt       string     `json:"-" gorm:"column:hash_salt"`
	HashedPassword string     `json:"-" gorm:"column:hashed_password"`
	VerifiedAt     *time.Time `json:"verified_at" gorm:"column:verified_at"`
	CreatedAt      time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Account) TableName() string {
	return "accounts"
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}
	return nil
}

func (a *Account) SetPassword(password string) (err error) {
	a.HashSalt = time.Now().Format(time.RFC3339Nano)
	saltedPassword := a.HashSalt + password
	hashedPasswordByte, err := bycrypt.GenerateFromPassword([]byte(saltedPassword), bycrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.HashedPassword = string(hashedPasswordByte)
	return nil
}

func (c Account) ComparePassword(password string) bool {
	saltedPassword := c.HashSalt + password
	return bycrypt.CompareHashAndPassword([]byte(c.HashedPassword), []byte(saltedPassword)) == nil
}

func GenerateRandomPassword() string {
	// generate random password of length 8
	long := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = long[rand.Int63()%int64(len(long))]
	}
	return string(b)
}

type Device struct {
	IP             string `json:"ip" gorm:"column:ip"`
	UserAgent      string `json:"user_agent" gorm:"column:user_agent"`
	Browser        string `json:"browser" gorm:"column:browser"`
	BrowserVersion string `json:"browser_version" gorm:"column:browser_version"`
	OS             string `json:"os" gorm:"column:os"`
	OSVersion      string `json:"os_version" gorm:"column:os_version"`
}

func (d Device) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

func (d Device) Value() (driver.Value, error) {
	return d.Marshal()
}

func (d *Device) Scan(src any) error {
	if src == nil {
		return nil
	}
	switch srcType := src.(type) {
	case []byte:
		return json.Unmarshal(srcType, d)
	case string:
		return json.Unmarshal([]byte(srcType), d)
	default:
		return errors.New("incompatible type for map")
	}
}

type Session struct {
	ID         string         `json:"id" gorm:"primaryKey;column:id"`
	AccountID  string         `json:"account_id" gorm:"column:account_id;not null"`
	Device     Device         `json:"device" gorm:"column:device;type:jsonb;"`
	Token      string         `json:"token" gorm:"column:token;not null"`
	ExpiresAt  time.Time      `json:"expires_at" gorm:"column:expires_at;not null"`
	LastActive time.Time      `json:"last_active" gorm:"column:last_active;not null;default:current_timestamp"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at;type:timestamptz"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index;null"`
}

func (Session) TableName() string {
	return "sessions"
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now().UTC()
	}
	return nil
}

func (s *Session) GenerateToken(signingKey string) error {
	if s.ExpiresAt.IsZero() {
		s.ExpiresAt = time.Now().Add(time.Hour * 1)
	}
	token, err := generateJWT(s.AccountID, signingKey, s.ExpiresAt)
	if err != nil {
		return err
	}
	s.Token = token
	return nil
}

func generateJWT(accountID, signingKey string, exp time.Time) (string, error) {
	claims := jwt.MapClaims{
		"account_id": accountID,
		"exp":        exp.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
