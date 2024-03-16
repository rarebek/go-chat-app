package router

import (
	"github.com/gin-gonic/gin"
	"go_chat_app/server/internal/user"
	"go_chat_app/server/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createroom", wsHandler.CreateRoom)
	r.GET("/ws/joinroom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getrooms", wsHandler.GetRooms)
	r.GET("/ws/getclients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
