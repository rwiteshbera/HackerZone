package routes

import (
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/controllers"
)

func AuthenticationRoutes(server *api.Server) {
	server.Router.POST("/signup", controllers.SignUp(server))
	server.Router.POST("/login", controllers.Login(server))
}
