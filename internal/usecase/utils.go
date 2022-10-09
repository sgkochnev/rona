package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	e "github.com/sgkochnev/rona/internal/error"
	"time"

	"github.com/sgkochnev/rona/internal/entity"
)

func createRefreshToken() (*entity.RefreshToken, error) {
	data := make([]byte, 32)

	if _, err := rand.Read(data); err != nil {
		return nil, err
	}

	token := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(token, data)

	rt := &entity.RefreshToken{
		ExpiresAt: time.Now().Add(expiresAtForRefresh).Unix(),
		IssuedAt:  time.Now().Unix(),
		Value:     string(token),
	}

	return rt, nil
}

func createAccessToken(user *entity.User, signedKey []byte) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAtForAccess)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: user.Username,
	})

	return accessToken.SignedString(signedKey)
}

func createToken(user *entity.User, signedKey []byte) (*entity.Token, error) {
	accessToken, err := createAccessToken(user, signedKey)
	if err != nil {
		return nil, e.ErrAccessToken
	}

	refreshToken, err := createRefreshToken()
	if err != nil {
		return nil, e.ErrRefreshToken
	}

	return &entity.Token{
		Access:       accessToken,
		RefreshToken: *refreshToken,
	}, nil
}
