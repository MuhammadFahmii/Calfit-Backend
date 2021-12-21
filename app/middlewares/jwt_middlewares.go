package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JWTMyClaims struct {
	Username string
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTMyClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(Username string) (string, error) {
	claims := JWTMyClaims{
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(jwtConf.ExpiresDuration)).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}
