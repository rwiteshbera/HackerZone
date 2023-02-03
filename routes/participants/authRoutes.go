package participants

import (
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/controllers/participants"
)

func AuthenticationRoutes(server *api.Server) {
	// Signup as a participant
	server.Router.POST("/signup/user", participants.SignUp(server))

	// Login as a participant
	server.Router.POST("/login/user", participants.Login(server))
}
