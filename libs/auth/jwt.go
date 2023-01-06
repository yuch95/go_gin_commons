package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtToken struct {
	secretKey   []byte
	Method      jwt.SigningMethod
	ExpiresTime time.Duration
}

func NewJwtToken(secretKey string, method jwt.SigningMethod, expiresTime time.Duration) *JwtToken {
	return &JwtToken{secretKey: []byte(secretKey), Method: method, ExpiresTime: expiresTime}
}

func (j JwtToken) GenerateToken(claim jwt.Claims) (string, error) {
	return jwt.NewWithClaims(j.Method, claim).SignedString(j.secretKey)
}

func (j JwtToken) ParseToken(tokenStr string, claim jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
