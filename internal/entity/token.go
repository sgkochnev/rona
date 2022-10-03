package entity

import (
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Access string
	RefreshToken
}

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type RefreshToken struct {
	IssuedAt  int64  `json:"issued_at" bson:"issued_at"`
	ExpiresAt int64  `json:"expires_at" bson:"expires_at"`
	Value     string `json:"token" bson:"token"`
}
