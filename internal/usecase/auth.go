package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sgkochnev/rona/internal/entity"
	e "github.com/sgkochnev/rona/internal/error"
	"golang.org/x/crypto/bcrypt"
)

var _ AuthUsecase = (*auth)(nil)

type auth struct {
	repo      AuthRepository
	signedKey []byte
}

func NewAuth(repo AuthRepository, signedKey []byte) *auth {
	return &auth{
		repo:      repo,
		signedKey: signedKey,
	}
}

func (a *auth) SignUp(ctx context.Context, user *entity.User) error {
	pwh, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error sign-in bcrypt.GenerateFromPassword: %w", err)
	}
	user.Password = string(pwh)

	return a.repo.Add(ctx, user)
}

func (a *auth) SignIn(ctx context.Context, user *entity.User) (*entity.Token, error) {
	userDB, err := a.repo.Get(ctx, user.Username)
	if err != nil {
		if errors.Is(err, e.ErrUserDoesNotExist) {
			return nil, e.ErrDataNotFound
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return nil, e.ErrInvalidPassword
	}

	token, err := createToken(user, a.signedKey)
	if err != nil {
		return nil, err
	}

	user.RefreshToken = token.RefreshToken
	err = a.repo.UpdateRefreshToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ParseToken
func (a *auth) ParseToken(accessToken string) (*entity.Claims, error) {
	claims := &entity.Claims{}
	_, err := jwt.ParseWithClaims(accessToken, claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, e.ErrSigningMethod
			}
			return a.signedKey, nil
		})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
