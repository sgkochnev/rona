package dto

import (
	"github.com/sgkochnev/rona/internal/entity"
	"net/mail"
	"strings"
)

const (
	minPasswordLength = 8
	numbers           = "0123456789"
	letters           = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols           = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

type UserAuthDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *UserAuthDTO) PasswordIsValid() bool {
	ok := len(u.Password) >= minPasswordLength
	return ok && strings.ContainsAny(u.Password, numbers) &&
		strings.ContainsAny(u.Password, uppercaseLetters) &&
		strings.ContainsAny(u.Password, symbols) &&
		strings.ContainsAny(u.Password, letters)
}

func (u *UserAuthDTO) IsValid() bool {
	return u.UsernameIsValid() &&
		u.PasswordIsValid()
}

func (u *UserAuthDTO) UsernameIsValid() bool {
	return len(u.Username) > 0
}

func (u *UserAuthDTO) User() *entity.User {
	user := &entity.User{
		Username: u.Username,
		Password: u.Password,
	}
	return user
}

type UserDTO struct {
	Email string `json:"email,omitempty"`
	UserAuthDTO
}

func (u *UserDTO) EmailIsValid() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}

func (u *UserDTO) IsValid() bool {
	return u.EmailIsValid() && u.UserAuthDTO.IsValid()
}

func (u *UserDTO) User() *entity.User {
	user := &entity.User{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
	return user
}
