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

func DataRoutes(server *api.Server) {
	// Get current logged in hackathon organizer's data
	server.Router.GET("/host", organizers.GetOrganizerData(server))
}

func HackathonRoutes(server *api.Server) {
	// Host a hackathon
	server.Router.POST("/host/create", organizers.CreateHack(server))

}
