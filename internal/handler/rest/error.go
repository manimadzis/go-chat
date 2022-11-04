package rest

import (
	"encoding/json"
	"net/http"
)

const (
	INTERNAL_SERVER_ERR = "Internal server error"
)

type errorDTO struct {
	error interface{} `json:"error"`
}

func (h *handler) sendError(w http.ResponseWriter, status int, data interface{}) {
	h.sendResponse(w, status, errorDTO{error: data})
}

func (h *handler) sendResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	jsonData, err := json.Marshal(data)
	if err != nil {
		h.logger.Error("Failed to Marshal errorDTO: %v: data=%v", err, data)
	}

	w.Write(jsonData)
}
