package tokens

import (
	"clean-architecture-service/config"
	"clean-architecture-service/internal/entity"
	"github.com/golang-jwt/jwt/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
)

var (
	JwtSignatureKey  = []byte(config.Cfg.App.JWTSecret)
	JwtExpireTime    = time.Now().Add(time.Hour * 24 * 7)
	JwtSigningMethod = jwt.SigningMethodHS256
)

type MyClaims struct {
	RegClaims jwt.RegisteredClaims
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
}

func (m MyClaims) Valid() error {
	return nil
}

func GenerateJWT(user *entity.User) (string, int64, error) {
	claims := MyClaims{
		RegClaims: jwt.RegisteredClaims{
			Issuer:    config.Cfg.App.Name,
			ExpiresAt: jwt.NewNumericDate(JwtExpireTime),
		},
		Name:     user.Name,
		LastName: user.LastName,
		ID:       user.ID.String(),
	}
	token := jwt.NewWithClaims(JwtSigningMethod, claims)
	signedJWT, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		return "", 0, err
	}
	return signedJWT, JwtExpireTime.Unix(), nil
}

func GenerateRefreshToken() (string, error) {
	id, err := gonanoid.New()
	return id, err
}
