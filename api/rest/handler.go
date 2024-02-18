package rest

import (
	fiber "github.com/gofiber/fiber/v2"

	"template-manager/internal/app/auth"
	"template-manager/internal/app/template"
	"template-manager/pkg/config"
)

type Middleware interface {
	FiberAuthMiddleware(c *fiber.Ctx) error
}

type server struct {
	conf        *config.Config
	authApp     *auth.App
	templateApp *template.App
	middleware  Middleware
}

// New creates a new fiber app
func New(
	conf *config.Config,
	authApp *auth.App,
	templateApp *template.App,
	middleware Middleware,
) *server {
	return &server{
		conf:        conf,
		authApp:     authApp,
		templateApp: templateApp,
		middleware:  middleware,
	}
}

func (s server) Listen(port string) error {
	app := fiber.New()

	app.Use(
		s.middleware.FiberAuthMiddleware,
	)

	// Setup route for the API health check
	app.Get("/health", health)
	app.Get("/stats", stats)

	api := app.Group("/api")

	// Define API endpoints for managing users
	api.Post("/users/signup", s.Signup)
	api.Post("/users/login", s.Login)
	api.Post("/users/reset-password", s.InitiateResetPassword)
	api.Post("/users/logout", s.Logout)

	// Define API endpoints for managing keys
	api.Post("/keys", s.AddKey)
	api.Get("/keys", s.ListAccessKeys)
	api.Delete("/keys/:id", s.DeleteKey)

	// Define API endpoints for managing templates
	api.Post("/templates/upload-url", s.GetUploadURL)
	api.Post("/templates", s.AddTemplate)
	api.Get("/templates/:id", s.GetTemplate)
	api.Get("/templates", s.ListTemplates)
	api.Put("/templates/:id", s.UpdateTemplate)
	api.Delete("/templates/:id", s.DeleteTemplate)

	// Start the server on port 8080
	return app.Listen(port)
}

func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func stats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"grpc":    false,
		"version": "v1.0.0",
		"open":    false, // open source version
	})
}

func HandleError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusUnprocessableEntity)
	return c.JSON(fiber.Map{
		"status":  false,
		"message": err.Error(),
	})
}

func HandleBadRequest(c *fiber.Ctx, message error) error {
	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Map{
		"status":  false,
		"message": message,
	})
}

func HandleSuccess(c *fiber.Ctx, message string, data any) error {
	r := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return c.JSON(r)
}
