package handlers

import (
	"CloudStorage/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type RenameFileRequest struct {
	NewFileName string `json:"new_file_name" validate:"required"`
}

func (h *Handler) RenameFile(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userID, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	var req RenameFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid request body"
		return
	}

	vars := mux.Vars(r)
	fileID, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid file ID"
		return
	}

	err = h.svc.RenameFile(fileID, int(userID), req.NewFileName)
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to rename file"
		return
	}

	resp = response.Success
	resp.Message = "File renamed successfully"
}
