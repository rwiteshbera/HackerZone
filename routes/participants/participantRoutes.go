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

func DataRoutes(server *api.Server) {
	// Get current logged in participant's data
	server.Router.GET("/user", participants.GetParticipantData(server))

	// Delete current logged in participant's data
	server.Router.DELETE("/user/delete", participants.DeleteParticipantData(server))

	// Join hackathon and create team
	server.Router.POST("/join", participants.JoinHack(server))

	// Get All Your Team Name
	server.Router.GET("/get_team", participants.GetAllTeamData(server)) // INCOMPLETE

	// Add Member in Team
	server.Router.POST("/add_member", participants.AddMemberInTeam(server))

	// Delete member from Team
	server.Router.DELETE("/delete_member", participants.DeleteMemberFromTeam(server))
}
