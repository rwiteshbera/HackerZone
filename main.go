package main

import (
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/config"
	"github.com/rwiteshbera/HackerZone/middlewares"
	"github.com/rwiteshbera/HackerZone/routes/organizers"
	"github.com/rwiteshbera/HackerZone/routes/participants"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config:", err)
	}

	Server, err := api.CreateServer(config)
	if err != nil {
		log.Fatalln("error:", err)
	}

	if config.SERVER_PORT == "" {
		config.SERVER_PORT = "5001"
	}

	participants.AuthenticationRoutes(Server)
	organizers.AuthenticationRoutes(Server)

	Server.Router.Use(middlewares.AuthMiddleware(Server))

	participants.DataRoutes(Server)
	organizers.DataRoutes(Server)
	organizers.HackathonRoutes(Server)

	err = Server.Start(config.SERVER_PORT)
	if err != nil {
		log.Fatalln("error:", err)
	}
}
