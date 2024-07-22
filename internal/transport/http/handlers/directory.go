package handlers

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/response"
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
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
