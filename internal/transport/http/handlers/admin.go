package handlers

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func (h *Handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	users, err := h.svc.AdminGetUserList()
	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = users
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	vars := mux.Vars(r)
	idStr := vars["id"]

	userId, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Invalid user ID"
		return
	}

	err = h.svc.DeleteUser(userId)
	if err != nil {
		if err == errors.ErrUserNotFound {
			resp = response.NotFound
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}

func (h *Handler) AdminGetUserList(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	users, err := h.svc.GetUserList()
	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = users
}

func (h *Handler) AdminUpdateUserById(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Invalid user ID"
		return
	}

	var inputData models.UserUpdateDto
	inputData.Id = id

	err = json.NewDecoder(r.Body).Decode(&inputData)
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

	err = h.svc.UpdateUser(inputData)
	if err != nil {
		if err == errors.ErrWrongPassword {
			resp.Code = http.StatusUnauthorized
			resp.Message = "Wrong current password"
			return
		} else if err == errors.ErrDataNotFound {
			resp = response.NotFound
			return
		} else {
			resp = response.InternalServer
			return
		}
	}

	resp = response.Success
}
