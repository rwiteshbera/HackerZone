package participants

import (
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/controllers/participants"
)

func AuthenticationRoutes(server *api.Server) {
	// Signup as a participant
	server.Router.POST("/signup", participants.SignUp(server))

	// Login as a participant
	server.Router.POST("/login", participants.Login(server))
}
