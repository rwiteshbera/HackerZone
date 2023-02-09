package organizers

import (
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/controllers/organizers"
)

func AuthenticationRoutes(server *api.Server) {
	// Signup as Hackathon organizer
	server.Router.POST("/signup/host", organizers.SignUpAsOrganizer(server))

	// Login as a Hackathon organizer
	server.Router.POST("/login/host", organizers.LoginAsOrganizer(server))
}
