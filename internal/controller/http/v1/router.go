package v1

import (
	"github.com/sgkochnev/rona/internal/controller/http/v1/middleware"
	"github.com/sgkochnev/rona/internal/controller/http/v1/utils"
	"github.com/sgkochnev/rona/internal/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sgkochnev/rona/pkg/logger"
)

// TODO add router and handlers

type Router struct {
	*mux.Router
	log   logger.Logger
	uCase Usecases
}

func NewRouter(l logger.Logger, uCase Usecases) *Router {
	r := &Router{
		Router: mux.NewRouter(),
		log:    l,
		uCase:  uCase,
	}

	r.endpoint()

	return r
}

func (r *Router) endpoint() {
	a := NewAuth(r.log, r.uCase)
	mdw := middleware.NewMdw(r.log, r.uCase)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", a.SignUp).Methods(http.MethodPost)
		auth.HandleFunc("/sign-in", a.SignIn).Methods(http.MethodPost)
	}
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.Use(mdw.UserIdentity)
	{
		v1.HandleFunc("/board", Board).Methods(http.MethodGet)
	}
}

// функция для проверки авторизации
func Board(w http.ResponseWriter, req *http.Request) {
	claims, ok := req.Context().Value(middleware.ContextKey).(*entity.Claims)
	if !ok {
		claims = &entity.Claims{}
	}

	log.Printf("%v", claims)

	w.Header().Set("Content-Type", "application/json")

	utils.ResponseJSON(w, http.StatusOK)
}
