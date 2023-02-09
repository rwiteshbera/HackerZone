package organizers

import (
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/controllers/organizers"
)

func DataRoutes(server *api.Server) {
	// Get current logged in hackathon organizer's data
	server.Router.GET("/host", organizers.GetOrganizerData(server))
}
