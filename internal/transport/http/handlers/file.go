package handlers

import (
	"CloudStorage/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userID, ok := context.Get(r, "user_id").(int64)
	if !ok {
		resp = response.Unauthorized
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		resp = response.BadRequest
		resp.Message = "Unable to parse form"
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		resp = response.BadRequest
		resp.Message = fmt.Sprintf("Unable to retrieve file from form: %v", err)
		return
	}
	defer file.Close()

	directoryIDStr := r.FormValue("directory_id")
	var directoryID int
	if directoryIDStr != "" {
		directoryID, err = strconv.Atoi(directoryIDStr)
		if err != nil {
			resp = response.BadRequest
			resp.Message = "Invalid directory_id"
			return
		}
	} else {
		directoryID = 0
	}

	err = h.svc.UploadFile(int(userID), directoryID, file, handler)
	if err != nil {
		resp = response.InternalServer
		resp.Message = fmt.Sprintf("Unable to upload file: %v", err)
		return
	}

	resp = response.Success
	resp.Message = "File uploaded successfully"
}

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
