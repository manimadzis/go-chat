package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-chat/internal/domain"
	"io"
	"net/http"
	"reflect"
)

func (h *handler) createChat(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.sendError(w, 500, INTERNAL_SERVER_ERR)
		h.logger.Errorf("Failed to read body: %v", err)
	}

	dto := domain.CreateChatDTO{}
	err = h.parseBytes(data, &dto)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err)
	}

	err = h.service.ChatService.Create(context.Background(), dto)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, INTERNAL_SERVER_ERR)
	}

	h.sendResponse(w, http.StatusNoContent, "")
}

func (h *handler) updateChat(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO: implement me
}

func (h *handler) deleteChat(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO: implement me
}

func (h *handler) getChat(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO: implement me
}

func (h *handler) parseBytes(data []byte, dto domain.DTO) error {
	err := json.Unmarshal(data, dto)
	if err != nil {
		h.logger.Errorf("Failed to parse %s: %v", reflect.TypeOf(dto).String(), err)
		return fmt.Errorf("parsing failed: %v", err)
	}
	err = dto.Valid()
	if err != nil {
		h.logger.Errorf("Validation of %s failed: %v", reflect.TypeOf(dto).String(), err)
		return fmt.Errorf("validation failed: %v", err)
	}
	return nil
}
