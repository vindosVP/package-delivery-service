package tokens

import (
	"clean-architecture-service/internal/entity"
	"github.com/golang-jwt/jwt/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
)

var (
	JwtSignatureKey  = []byte("jwtSecret")
	JwtSigningMethod = jwt.SigningMethodHS256
)

type MyClaims struct {
	RegisteredClaims jwt.RegisteredClaims
	ID               string `json:"id"`
	Name             string `json:"name"`
	LastName         string `json:"lastName"`
}

func (m MyClaims) Valid() error {
	return nil
}

func GenerateJWT(user *entity.User) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24 * 7)
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "orders-service",
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		ID:   user.ID.String(),
		Name: user.Name,
	}

	token := jwt.NewWithClaims(JwtSigningMethod, claims)
	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		return "", 0, err
	}

	return signedToken, exp.Unix(), nil
}

func GenerateRefreshToken() (string, error) {
	id, err := gonanoid.New()
	return id, err
}
