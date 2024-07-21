package handlers

import (
	"CloudStorage/pkg/response"
	"net/http"
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
