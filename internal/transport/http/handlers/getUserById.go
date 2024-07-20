package handlers

import (
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
