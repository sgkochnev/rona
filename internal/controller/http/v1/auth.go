package v1

import (
	"errors"
	"github.com/sgkochnev/rona/internal/controller/http/v1/dto"
	"github.com/sgkochnev/rona/internal/controller/http/v1/utils"
	e "github.com/sgkochnev/rona/internal/error"
	"github.com/sgkochnev/rona/pkg/logger"
	"net/http"
)

type auth struct {
	log   logger.Logger
	uCase Usecases
}

func NewAuth(log logger.Logger, uCase Usecases) *auth {
	return &auth{
		log:   log,
		uCase: uCase,
	}
}

func (a *auth) SignUp(w http.ResponseWriter, r *http.Request) {

	var userDTO dto.UserDTO

	user, err := utils.ValidateUserData(r, &userDTO)
	if err != nil {
		a.log.Error("%v", err)
		utils.ResponseJSONWithErr(w, http.StatusBadRequest)
		return
	}

	err = a.uCase.SignUp(r.Context(), user)
	if err != nil {
		a.log.Error("error: %v", err.Error())
		utils.ResponseJSONWithErr(w, http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated)
}

func (a *auth) SignIn(w http.ResponseWriter, r *http.Request) {

	var userAuthDTO dto.UserAuthDTO

	user, err := utils.ValidateUserData(r, &userAuthDTO)
	if err != nil {
		a.log.Error("error validation: %v", err)
		utils.ResponseJSONWithErr(w, http.StatusBadRequest)
		return
	}

	token, err := a.uCase.SignIn(r.Context(), user)
	if err != nil {
		a.log.Error("error SignIn: %v", err.Error())

		switch {
		case errors.Is(err, e.ErrInvalidPassword),
			errors.Is(err, e.ErrDataNotFound):
			utils.ResponseJSONWithErr(w, http.StatusBadRequest)
		default:
			utils.ResponseJSONWithErr(w, http.StatusInternalServerError)
		}
		return
	}

	utils.WriteToken(w, token)

	utils.ResponseJSON(w, http.StatusOK)
}
