package handlers

import (
	"CloudStorage/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
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
