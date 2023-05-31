package database

import (
	"clean-architecture-service/config"
	"clean-architecture-service/internal/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGorm(cfg config.DB) (*gorm.DB, error) {
	dns := GenerateGormDNS(cfg)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database - NewGorm - gorm.Open: %w", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.Token{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.Package{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.Delivery{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	DB = db

	return db, err
}

func GenerateGormDNS(cfg config.DB) string {
	if cfg.DNS != "" {
		return cfg.DNS
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name, cfg.SSLMode)
}
