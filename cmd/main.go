package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"strings"
	"template-manager/internal/app"
	"template-manager/internal/entity"

	"template-manager/api/middleware"
	"template-manager/api/rest"
	"template-manager/internal/app/session"
	"template-manager/internal/pkg/email/mailjet"
	"template-manager/pkg/config"
	"template-manager/pkg/database"
	"template-manager/pkg/repository"
)

func cleanPort(port string) string {
	if strings.Contains(port, ":") {
		return port
	}
	return ":" + port
}

func main() {
	var server, port string
	flag.StringVar(&server, "server", "rest", "grpc or rest")
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.Parse()
	port = cleanPort(port)

	conf := loadConfig()

	db, err := database.New(
		conf.GetString("POSTGRES_DSN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: migrated this to repository.NewRepository()
	err = db.Client.AutoMigrate(
		&entity.Account{},
		&entity.Key{},
		&entity.Session{},
		&entity.Template{},
		&entity.TemplateSync{},
	)
	if err != nil {
		log.Fatal(err)
	}
	mj := mailjet.New(
		conf.GetString("MAILJET_PUBLIC_KEY"),
		conf.GetString("MAILJET_PRIVATE_KEY"),
		conf.GetString("MAILJET_DEFAULT_SENDER"),
		mailjet.WithName("template manager"),
	)
	logger := slog.Default()
	sessionManager := session.New(db.Client, conf, logger)
	midware := middleware.NewAuth(sessionManager)
	repo := repository.NewRepositoryContainer(db)

	apps := app.NewApp(conf, mj, logger, repo, sessionManager)

	restApp := rest.New(
		conf,
		apps.AuthApp,
		apps.TemplateApp,
		midware,
	)
	log.Fatal(restApp.Listen(port))
}

func loadConfig() *config.Config {
	conf := config.New().
		SetEnv("MAILJET_DOMAIN", os.Getenv("MAILJET_DOMAIN")).
		SetEnv("MAILJET_PRIVATE_KEY", os.Getenv("MAILJET_PRIVATE_KEY")).
		SetEnv("MAILJET_PUBLIC_KEY", os.Getenv("MAILJET_PUBLIC_KEY")).
		SetEnv("MAILJET_DEFAULT_SENDER", os.Getenv("MAILJET_DEFAULT_SENDER")).
		SetEnv("POSTGRES_DSN", os.Getenv("POSTGRES_DSN")).
		SetEnv("JWT_SIGNING_KEY", os.Getenv("JWT_SIGNING_KEY"))
	return conf
}
