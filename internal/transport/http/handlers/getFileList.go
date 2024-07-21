package handlers

import (
	"CloudStorage/pkg/response"
	"net/http"

	"github.com/gorilla/context"
)

func (h *Handler) GetFileList(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userId, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	files, err := h.svc.GetFileList(int(userId))

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = files
}
