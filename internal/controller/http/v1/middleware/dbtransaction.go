package middleware

import (
	"clean-architecture-service/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

func GormTransaction(db *gorm.DB, l logger.Interface) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tx := db.Begin()
		l.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		c.Locals("db_trx", tx)
		err := c.Next()
		if err != nil {
			l.Error(err, "middleware - Transaction - c.Next")
			l.Info("rolling back transaction due to error")
			tx.Rollback()
		}

		if StatusInList(c.Response().StatusCode(), []int{http.StatusOK, http.StatusCreated}) {
			l.Info("committing transactions")
			if err := tx.Commit().Error; err != nil {
				l.Error(err, "trx commit error")
			}
		} else {
			l.Info(fmt.Sprintf("rolling back transaction due to status code: %d", c.Response().StatusCode()))
			tx.Rollback()
		}
		return nil
	}
}

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}
