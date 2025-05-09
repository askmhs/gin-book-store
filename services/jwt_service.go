package services

import (
	"fmt"
	"time"

	"github.com/askmhs/gin-book-store/config"
	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() *JwtService {
	config.LoadConfig()
	secretKey := config.AppConfig.JwtSecret
	jwtIssuer := config.AppConfig.JwtIssuer

	return &JwtService{
		secretKey: secretKey,
		issuer:    jwtIssuer,
	}
}

func (s *JwtService) GenerateToken(data map[string]any) (string, error) {
	claims := &jwtCustomClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			Issuer:    config.AppConfig.JwtIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JwtService) ValidateToken(encodedToken string) (map[string]any, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (any, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
