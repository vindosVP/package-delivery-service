package token_repo

import (
	"clean-architecture-service/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

func (t *TokenRepo) Create(UserID uuid.UUID, token string) error {
	return t.db.Create(&entity.Token{UserID: UserID, Token: token}).Error
}

func (t *TokenRepo) GetActiveToken(UserID uuid.UUID) (*entity.Token, error) {
	var token entity.Token
	err := t.db.Where("user_id = ? AND revoked = ?", UserID, false).First(&token).Error
	return &token, err
}

func (t *TokenRepo) Revoke(UserID uuid.UUID) error {
	return t.db.Model(&entity.Token{}).Where("user_id = ?", UserID).Update("revoked", true).Error
}
