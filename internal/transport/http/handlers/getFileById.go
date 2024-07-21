package handlers

import (
	"CloudStorage/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func (h *Handler) GetFileById(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Invalid file ID"
		return
	}

	userId, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}
	file, err := h.svc.GetFileById(id, int(userId))

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = file
}
