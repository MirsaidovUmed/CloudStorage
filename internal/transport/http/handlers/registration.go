package handlers

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/response"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var inputData models.User

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		resp = response.BadRequest
		return
	}

	validate := validator.New()

	err = validate.Struct(inputData)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Invalid input data"
		return
	}

	err = h.svc.Registration(inputData)
	if err != nil {
		if err == errors.ErrAlreadyHasUser {
			resp.Code = 409
			resp.Message = "Пользователь с таким email уже существует"
		} else if err == errors.ErrDataNotFound {
			resp = response.BadRequest
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
