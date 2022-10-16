package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/sgkochnev/rona/internal/controller/http/v1/dto"
	"github.com/sgkochnev/rona/internal/entity"
	e "github.com/sgkochnev/rona/internal/error"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func response(w http.ResponseWriter, code int, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(js); err != nil {
		log.Println(err)
	}

}

func generateResponse(code int) Response {
	r := Response{Code: code}
	switch code {
	case http.StatusOK:
		r.Message = "OK"
	case http.StatusCreated:
		r.Message = "Created"
	case http.StatusMovedPermanently:
		r.Message = "Moved"
	}
	return r
}

func RJSON(w http.ResponseWriter, code int, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response(w, code, js)
}

func ResponseJSON(w http.ResponseWriter, code int) {
	RJSON(w, code, generateResponse(code))
}

func ResponseJSONWithErr(w http.ResponseWriter, code int) {
	RJSON(w, code, e.HTTPResponseError(code))
}

func WriteToken(w http.ResponseWriter, token *entity.Token) {
	cookie := &http.Cookie{
		Name:     "RefreshToken",
		HttpOnly: true,
		Secure:   true,
		Value:    token.RefreshToken.Value,
		Expires:  time.Now().AddDate(0, 0, 14),
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.Access))
}

type userDTO interface {
	*dto.UserDTO | *dto.UserAuthDTO
	IsValid() bool
	User() *entity.User
}

// ValidateUserData сhecks data received by http. Returns user and error
func ValidateUserData[T userDTO](r *http.Request, u T) (*entity.User, error) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		return nil, fmt.Errorf("json decoding failed: %w", err)
	}

	if !u.IsValid() {
		return nil, e.ErrUserData
	}

	user := u.User()
	user.Fingerprint = fingerprint(r, user.Username)

	return user, nil
}

func fingerprint(r *http.Request, salt string) string {
	accept := r.Header.Get("Accept")
	acceptLanguage := r.Header.Get("Accept-Language")
	secChUa := r.Header.Get("sec-ch-ua")
	upgradeInsecureRequests := r.Header.Get("Upgrade-Insecure-Requests")
	userAgent := r.Header.Get("User-Agent")

	fingerprint := fmt.Sprintf("%s:%s:%s:%s:%s",
		accept, acceptLanguage, secChUa, upgradeInsecureRequests, userAgent)

	h := sha256.New()
	h.Write([]byte(fingerprint + salt))

	return fmt.Sprintf("%x", h.Sum(nil))
}
