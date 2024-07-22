package handlers

import (
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/response"
	"net/http"
	"strconv"

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
