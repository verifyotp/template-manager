package session

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"

	"template-manager/internal/entity"
	"template-manager/pkg/config"
)

type Session struct {
	db     *gorm.DB
	config *config.Config
	logger *slog.Logger
}

func New(db *gorm.DB, conf *config.Config, logger *slog.Logger) *Session {
	return &Session{
		db:     db,
		config: conf,
		logger: logger,
	}
}

func (s *Session) Create(ctx context.Context, accountID string, device entity.Device) (*entity.Session, error) {

	// create session
	var sess = entity.Session{
		AccountID: accountID,
		Device:    device,
		ExpiresAt: time.Now().Add(time.Hour * 2),
	}
	err := sess.GenerateToken(s.config.GetString("JWT_SIGNING_KEY"))
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to generate token %+v", err)
		return nil, err
	}
	if err := s.db.Model(&sess).Create(&sess).Error; err != nil {
		s.logger.ErrorContext(ctx, "failed to create session %+v", err)
		return nil, err
	}

	return &sess, nil
}

func (s *Session) Verify(ctx context.Context, token string) (*entity.Session, error) {

	// extract account id from token
	jwtClaims, err := extractClaims(token, s.config.GetString("JWT_SIGNING_KEY"))
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to extract claims %+v", err)
		return nil, err
	}

	var sess entity.Session
	if err := s.db.Model(&sess).
		Where("token = ? AND account_id = ?", token, jwtClaims["account_id"]).
		First(&sess).Error; err != nil {
		s.logger.ErrorContext(ctx, "failed to find session %+v", err)
		return nil, err
	}
	// check if session is expired
	if sess.ExpiresAt.Before(time.Now()) {
		// delete session
		if err := s.db.Model(&sess).Delete(&sess).Error; err != nil {
			s.logger.ErrorContext(ctx, "failed to delete session %+v", err)
			return nil, err
		}
		return nil, errors.New("session expired")
	}

	return &sess, nil
}

func extractClaims(token, signingKey string) (map[string]any, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	// check if token is valid
	if !parsed.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

// Expire expires a session
func (s *Session) Expire(ctx context.Context, token string) error {
	var sess entity.Session
	if err := s.db.Model(&sess).Where("token = ?", token).First(&sess).Error; err != nil {
		s.logger.ErrorContext(ctx, "failed to find session %+v", err)
		return err
	}
	// delete session
	if err := s.db.Model(&sess).Delete(&sess).Error; err != nil {
		s.logger.ErrorContext(ctx, "failed to delete session %+v", err)
		return err
	}
	return nil
}

// Delete deletes all sessions for a given account
func (s *Session) Delete(ctx context.Context, accountID string) error {
	var sess entity.Session
	if err := s.db.Model(&sess).Where("account_id = ?", accountID).First(&sess).Error; err != nil {
		s.logger.ErrorContext(ctx, "failed to find session %+v", err)
		return err
	}
	// delete session
	if err := s.db.Model(&sess).Delete(&sess).Error; err != nil {
		s.logger.ErrorContext(ctx, "failed to delete session %+v", err)
		return err
	}
	return nil
}
