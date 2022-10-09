package middleware

import (
	"context"
	"github.com/sgkochnev/rona/internal/controller/http/v1/utils"
	"github.com/sgkochnev/rona/internal/entity"
	"github.com/sgkochnev/rona/pkg/logger"
	"net/http"
	"strings"
)

const ContextKey = "claims"

type TokenParser interface {
	ParseToken(string) (*entity.Claims, error)
}

type mdw struct {
	l     logger.Logger
	uCase TokenParser
}

func NewMdw(l logger.Logger, uCase TokenParser) *mdw {
	return &mdw{
		l:     l,
		uCase: uCase,
	}
}

func (m *mdw) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			m.l.Info("empty auth header")
			utils.ResponseJSONWithErr(w, http.StatusUnauthorized)
			return
		}

		header = strings.TrimPrefix(header, "Bearer ")

		claims, err := m.uCase.ParseToken(header)
		if err != nil {
			m.l.Info("invalid token: %v", err)
			utils.ResponseJSONWithErr(w, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
