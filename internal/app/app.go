package app

import (
	"clean-architecture-service/config"
	v1 "clean-architecture-service/internal/controller/http/v1"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/internal/usecase/package_repo"
	"clean-architecture-service/internal/usecase/token_repo"
	"clean-architecture-service/internal/usecase/user_repo"
	"clean-architecture-service/internal/validations"
	"clean-architecture-service/pkg/database"
	"clean-architecture-service/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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

	userUseCase := usecase.NewUserUseCase(user_repo.New(db), token_repo.New(db))
	packageUseCase := usecase.NewPackageUseCase(package_repo.New(db))
	handler := fiber.New()

	handler.Use(fiberlog.New(fiberlog.Config{
		TimeZone: "Europe/Moscow",
		Format:   "[${time}] ${locals:request-id} ${status} - ${latency} ${method} ${path}​\n",
	}))

	handler.Use(requestid.New(requestid.Config{
		Header:     "X-Request-ID",
		ContextKey: "request-id",
	}))

	v1.SetupRouter(handler, userUseCase, packageUseCase, lg)

	if err := handler.Listen(cfg.App.Port); err != nil {
		lg.Fatal(fmt.Errorf("app - Run - handler.Listen: %w", err))
	}
}
