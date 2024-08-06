package main

import (
	"log"
	"macbookpro/realtime-chat/server/database"
	"macbookpro/realtime-chat/server/internal/user"
	"macbookpro/realtime-chat/server/internal/ws"
	"macbookpro/realtime-chat/server/router"
)

func main() {
	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
