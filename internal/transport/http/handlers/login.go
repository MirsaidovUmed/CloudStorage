package handlers

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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
		resp.Message = fmt.Sprintf("Invalid input data : %v", err)
		return
	}

	token, err := h.svc.Login(inputData)

	if err != nil {
		if err == errors.ErrDataNotFound {
			resp = response.NotFound
			return
		} else if err == errors.ErrWrongPassword {
			resp.Code = 401
			resp.Message = "Wrong Password"
			return
		} else {
			resp = response.InternalServer
			return
		}
	}

	resp = response.Success
	resp.Payload = token
}
