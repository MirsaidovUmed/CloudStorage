package handlers

import (
	"CloudStorage/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func (h *Handler) ShareFile(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userId, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	vars := mux.Vars(r)
	fileId, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid file ID"
		return
	}

	targetUserId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid user ID"
		return
	}

	err = h.svc.ShareFile(int(userId), fileId, targetUserId)
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to share file"
		return
	}

	resp = response.Success
	resp.Message = "File shared successfully"
}

func (h *Handler) GetFileAccessUsers(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	vars := mux.Vars(r)
	fileId, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid file ID"
		return
	}

	users, err := h.svc.GetFileAccessUsers(fileId)
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to get file access users"
		return
	}

	resp = response.Success
	resp.Payload = users
}
