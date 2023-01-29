package routes

import (
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/controllers"
	"github.com/rwiteshbera/HackZone/middlewares"
)

func DataRoutes(server *api.Server) {
	server.Router.Use(middlewares.AuthMiddleware(server))
	server.Router.GET("/user", controllers.GetData(server))
}
