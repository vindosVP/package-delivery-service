package tokens

import (
	"clean-architecture-service/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"strings"
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

type TokenMetadata struct {
	UserID  uuid.UUID
	Expires int64
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			UserID:  userID,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return JwtSignatureKey, nil
}
