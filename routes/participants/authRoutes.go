package participants

import (
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/controllers/participants"
)

func AuthenticationRoutes(server *api.Server) {
	// Signup as a participant
	server.Router.POST("/signup/user", participants.SignUp(server))

	// Login as a participant
	server.Router.POST("/login/user", participants.Login(server))
}
