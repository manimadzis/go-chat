package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]

func (h *handler) signUp(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//json.Unmarshal()r.URL.Query().
}

func (h *handler) signIn(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO: implement me
}
