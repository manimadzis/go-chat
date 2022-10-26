package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	router *httprouter.Router
}

func (h *handler) initRouter() {
	h.router.POST("/signup", h.signUp)
	h.router.GET("/signin", h.signUp)

	h.router.POST("/msg/", h.sendMessage)
	h.router.GET("/msg/:id", h.getMessage)
	h.router.PUT("/msg/:id", h.updateMessage)
	h.router.DELETE("/msg/:id", h.deleteMessage)

	h.router.POST("/chat/", h.createChat)
	h.router.GET("/chat/:id", h.getChat)
	h.router.PUT("/chat/:id", h.updateChat)
	h.router.DELETE("/chat/:id", h.deleteChat)

	h.router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func New(router *httprouter.Router) http.Handler {
	h := handler{router: router}
	h.initRouter()
	return &h
}
