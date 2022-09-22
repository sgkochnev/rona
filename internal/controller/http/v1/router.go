package v1

import (
	"net/http"

	"github.com/sgkochnev/rona/pkg/logger"
)

// TODO add router and handlers

type usecases interface {
}

func NewRouter(handler *http.ServeMux, l logger.Logger, uCase usecases) {
}
