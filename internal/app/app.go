package app

import (
	"clean-architecture-service/config"
	v1 "clean-architecture-service/internal/controller/http/v1"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/internal/usecase/repo"
	"clean-architecture-service/internal/validations"
	"clean-architecture-service/pkg/database"
	"clean-architecture-service/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) {
	lg := logger.New(cfg.Log.Level)

	db, err := database.NewGorm(cfg.DB)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - database.NewGorm: %w", err))
	}

	if err := validations.InitValidations(); err != nil {
		lg.Fatal(fmt.Errorf("app - Run - validations.InitValidations: %w", err))
	}

	userUseCase := usecase.New(repo.New(db))
	handler := fiber.New()
	v1.SetupRouter(handler, userUseCase, lg)

	if err := handler.Listen(cfg.App.Port); err != nil {
		lg.Fatal(fmt.Errorf("app - Run - handler.Listen: %w", err))
	}
}
