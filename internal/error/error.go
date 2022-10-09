package error

import (
	"errors"
	"net/http"
)

// The errors we want to use inside the controller.
var (
	ErrBadRequest          = errors.New("bad request")           // 400
	ErrUnauthorized        = errors.New("unauthorized")          // 401
	ErrForbidden           = errors.New("forbidden access")      // 403
	ErrNotFound            = errors.New("resource not found")    // 404
	ErrNotAllowed          = errors.New("operation not allowed") // 405
	ErrConflict            = errors.New("datamodel conflict")    // 409
	ErrGone                = errors.New("resource gone")         // 410
	ErrUnprocessableEntity = errors.New("unprocessable entity")  // 422
	ErrInternalServer      = errors.New("internal server error") // 500

)

// The errors we want to use inside the business logic.
var (
	ErrUserData        = errors.New("invalid user data")
	ErrInvalidPassword = errors.New("invalid password")
	ErrAccessToken     = errors.New("can not create access token")
	ErrRefreshToken    = errors.New("can not create refresh token")
	ErrSigningMethod   = errors.New("invalid signing method")
	ErrDataNotFound    = errors.New("data not found")
)

// The errors we want to use inside the repository
var (
	ErrDuplicateEntry   = errors.New("duplicate entry")
	ErrUserDoesNotExist = errors.New("user does not exist")
)

// HTTPError is our custom HTTP error to get a proper string output.
type HTTPError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func HTTPResponseError(code int) HTTPError {
	he := HTTPError{Code: code}
	switch code {
	case http.StatusBadRequest:
		he.Message = ErrBadRequest.Error()
	case http.StatusUnauthorized:
		he.Message = ErrUnauthorized.Error()
	case http.StatusForbidden:
		he.Message = ErrForbidden.Error()
	case http.StatusNotFound:
		he.Message = ErrNotFound.Error()
	case http.StatusMethodNotAllowed:
		he.Message = ErrNotAllowed.Error()
	case http.StatusConflict:
		he.Message = ErrConflict.Error()
	case http.StatusGone:
		he.Message = ErrGone.Error()
	case http.StatusUnprocessableEntity:
		he.Message = ErrUnprocessableEntity.Error()
	default:
		he.Code = http.StatusInternalServerError
		he.Message = ErrInternalServer.Error()
	}
	return he
}
