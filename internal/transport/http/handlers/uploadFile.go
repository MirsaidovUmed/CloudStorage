package handlers

import (
	"CloudStorage/pkg/response"
	"fmt"
	"net/http"

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
		resp.Message = fmt.Sprintf("Unable to retrieve file from form %v", err)
		return
	}
	defer file.Close()

	err = h.svc.UploadFile(int(userID), file, handler)
	if err != nil {
		resp = response.InternalServer
		resp.Message = fmt.Sprintf("Unable to upload file: %v", err)
		return
	}

	resp = response.Success
	resp.Message = "File uploaded successfully"
}
