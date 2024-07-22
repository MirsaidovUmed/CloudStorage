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

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var inputData models.UserCreateDto

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

	err = h.svc.Registration(inputData)
	if err != nil {
		if err == errors.ErrAlreadyHasUser {
			resp.Code = 409
			resp.Message = "Пользователь с таким email уже существует"
			return
		} else if err == errors.ErrDataNotFound {
			resp = response.BadRequest
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var inputData models.UserUpdateDto

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

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.svc.GetUserByID(id)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp = response.NotFound
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = user
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.UserCreateDto

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
