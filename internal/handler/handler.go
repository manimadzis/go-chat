package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h Handler) InitRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/signup", h.signUp)
		user.POST("/signin", h.signIn)
	}
	msg := router.Group("/msg")
	{
		msg.POST("/", h.sendMessage)
		msg.GET("/:id", h.getMessage)
		msg.PUT("/:id", h.updateMessage)
		msg.DELETE("/:id", h.deleteMessage)
	}
	chat := router.Group("/chat")
	{
		chat.POST("/", h.createChat)
		chat.GET("/:id", h.getChat)
		chat.PUT("/:id", h.updateChat)
		chat.DELETE("/:id", h.deleteChat)
	}
}
