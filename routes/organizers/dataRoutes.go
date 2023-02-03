package organizers

import (
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/controllers/organizers"
)

func DataRoutes(server *api.Server) {
	// Get current logged in hackathon organizer's data
	server.Router.GET("/host", organizers.GetOrganizerData(server))
}
