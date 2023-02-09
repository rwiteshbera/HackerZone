package participants

import (
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/controllers/participants"
)

func DataRoutes(server *api.Server) {
	// Get current logged in participant's data
	server.Router.GET("/user", participants.GetParticipantData(server))

	// Delete current logged in participant's data
	server.Router.DELETE("/user/delete", participants.DeleteParticipantData(server))
}
