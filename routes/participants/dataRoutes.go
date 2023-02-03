package participants

import (
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/controllers/participants"
)

func DataRoutes(server *api.Server) {
	// Get current logged in participant's data
	server.Router.GET("/user", participants.GetParticipantData(server))

	// Delete current logged in participant's data
	server.Router.DELETE("/user/delete", participants.DeleteParticipantData(server))
}
