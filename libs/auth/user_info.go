package auth

import "github.com/golang-jwt/jwt/v4"

type UserInfo struct {
	UserGuid int
	UserName string
}

type MyClaim struct {
	*UserInfo
	jwt.RegisteredClaims
}
