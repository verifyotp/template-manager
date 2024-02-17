package middleware

import (
	"context"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"template-manager/internal/entity"
)

type Auth struct {
	sess SessionManager
}

func NewAuth(sess SessionManager) *Auth {
	return &Auth{
		sess: sess,
	}
}

type SessionManager interface {
	Verify(ctx context.Context, token string) (*entity.Session, error)
}

var unauthenticatedRoutes = map[string]bool{
	"/stats":                    true,
	"/health":                   true,
	"/api/users/signup":         true,
	"/api/users/login":          true,
	"/api/users/reset-password": true,
}

func (a *Auth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if _, ok := unauthenticatedRoutes[r.URL.Path]; ok {
			next.ServeHTTP(w, r)
			return
		}

		// extract token from header
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// verify token
		sess, err := a.sess.Verify(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// set account id in context
		ctx := context.WithValue(r.Context(), "account_id", sess.AccountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Auth) FiberAuthMiddleware(c *fiber.Ctx) error {

	path := c.Path()
	if _, ok := unauthenticatedRoutes[path]; ok {
		return c.Next()
	}

	// extract token from header
	token := c.Get("Authorization")
	if token == "" {
		return c.SendStatus(http.StatusUnauthorized)
	}

	// verify token
	sess, err := a.sess.Verify(c.Context(), token)
	if err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	// set account id in context
	ctx := context.WithValue(c.Context(), "account_id", sess.AccountID)
	c.SetUserContext(ctx)
	c.Context().SetUserValue("account_id", sess.AccountID)

	return c.Next()
}
