package app

import (
	"clean-architecture-service/config"
	"clean-architecture-service/pkg/database"
	"clean-architecture-service/pkg/logger"
	"fmt"
)

func Run(cfg *config.Config) {
	lg := logger.New(cfg.Log.Level)

	dns := GenerateGormDNS(cfg.DB)
	_, err := database.NewGorm(dns)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - database.NewGorm: %w", err))
	}
}

func GenerateGormDNS(cfg config.DB) string {
	if cfg.DNS != "" {
		return cfg.DNS
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name, cfg.SSLMode)
}
