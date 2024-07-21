package handlers

import (
	"CloudStorage/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func (h *Handler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userID, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	vars := mux.Vars(r)
	fileID, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid file ID"
		return
	}

	err = h.svc.RemoveFile(fileID, int(userID))
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to delete file"
		return
	}

	resp = response.Success
	resp.Message = "File deleted successfully"
}
