package handlers

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateDirectory(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var inputData models.Directory

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid input data"
		return
	}

	userID, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		resp.Message = "Unauthorized"
		return
	}
	inputData.UserId = int(userID)

	err = h.svc.CreateDirectory(inputData)
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to create directory"
		return
	}

	resp = response.Success
	resp.Message = "Directory created successfully"
}

type RenameDirRequest struct {
	NewFileName string `json:"new_dir_name" validate:"required"`
}

func (h *Handler) RenameDirectory(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userID, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	var req RenameDirRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid request body"
		return
	}

	vars := mux.Vars(r)
	dirId, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid directory ID"
		return
	}

	err = h.svc.RenameDirectory(dirId, int(userID), req.NewFileName)
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to rename directory"
		return
	}

	resp = response.Success
	resp.Message = "Directory renamed successfully"
}

func (h *Handler) GetDirectoryById(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Invalid directory ID"
		return
	}

	userId, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}
	dir, err := h.svc.GetDirectoryById(id, int(userId))

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = dir
}

func (h *Handler) DeleteDirectory(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userId, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	vars := mux.Vars(r)
	dirId, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Invalid dir Id"
		return
	}

	err = h.svc.DeleteDirectory(dirId, int(userId))
	if err != nil {
		resp = response.InternalServer
		resp.Message = "Unable to delete directory"
		return
	}

	resp = response.Success
	resp.Message = "Directory deleted successfully"
}
