package main

import (
	"go_chat_app/server/db"
	"go_chat_app/server/internal/user"
	"go_chat_app/server/internal/ws"
	"go_chat_app/server/router"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userResp := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userResp)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("localhost:8080")
}
