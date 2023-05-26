package database

import (
	"clean-architecture-service/internal/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func NewGorm(dns string) (*GormDB, error) {

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database - NewGorm - gorm.Open: %w", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("database - NewGorm - db.AutoMigrate: %w", err)
	}

	return &GormDB{db: db}, err
}
