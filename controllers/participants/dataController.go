package participants

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/controllers"
	"github.com/rwiteshbera/HackerZone/database"
	"github.com/rwiteshbera/HackerZone/models"
	"net/http"
	"strings"
)

// GetParticipantData : Get logged in participants data
func GetParticipantData(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		currentEmail, err := context.Cookie("email")
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err := database.Connect(server)
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
				return
			}
		}(db)

		response, isExist := GetUserDataByEmail(db, currentEmail)
		if !isExist {
			controllers.LogErrorWithAbort(context, errors.New("unable to find participant"), http.StatusInternalServerError)
			return
		}

		var useDataResponse = &models.User{
			FirstName: response.FirstName,
			LastName:  response.LastName,
			Email:     response.Email,
			Bio:       response.Bio,
			Gender:    response.Gender,
			LastLogin: response.LastLogin,
			CreatedAt: response.CreatedAt,
		}

		context.JSON(http.StatusOK, gin.H{"message": useDataResponse})
	}
}

// DeleteParticipantData : Delete Logged In participants data
func DeleteParticipantData(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		currentEmail, err := context.Cookie("email")
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err := database.Connect(server)
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
				return
			}
		}(db)

		// Delete participant data based on current logged in email
		_, err = db.Exec(database.DeleteParticipantQuery, currentEmail)
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		controllers.SendResponse(context, "You profile has been delete successfully")
	}
}

// JoinHack : Join hackathon and create team
func JoinHack(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var teamData models.Team

		err1 := context.ShouldBindJSON(&teamData)
		if err1 != nil {
			controllers.LogErrorWithAbort(context, err1, http.StatusBadRequest)
			return
		}

		currentUserId, err := context.Cookie("user")
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		currentUserEmail, err2 := context.Cookie("email")
		if err2 != nil {
			controllers.LogErrorWithAbort(context, err2, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err3 := database.Connect(server)
		if err3 != nil {
			controllers.LogErrorWithAbort(context, err3, http.StatusInternalServerError)
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			}
		}(db)

		teamId, err4 := uuid.NewUUID()
		if err4 != nil {
			controllers.LogErrorWithAbort(context, err4, http.StatusInternalServerError)
			return
		}

		teamData.TeamId = teamId

		// Upload team and join hackathon
		_, err5 := db.Exec(database.CreateTeamQuery, teamData.TeamId, teamData.HackathonId, teamData.Name, currentUserId)
		if err5 != nil {
			if strings.Contains(err5.Error(), "foreign key constraint") {
				controllers.LogErrorWithAbort(context, errors.New("team name is not available"), http.StatusInternalServerError)
				return
			} else if strings.Contains(err5.Error(), "duplicate key") {
				controllers.LogErrorWithAbort(context, errors.New("team name is not available"), http.StatusInternalServerError)
				return
			} else {
				controllers.LogErrorWithAbort(context, errors.New("something went wrong"), http.StatusInternalServerError)
				return
			}
		}

		// Add yourself in the team also
		_, err6 := db.Exec(database.InsertMemberInTeamQuery, teamData.TeamId, currentUserEmail, teamData.HackathonId)
		if err6 != nil {
			controllers.LogErrorWithAbort(context, errors.New("something went wrong"), http.StatusInternalServerError)
			return
		}

		controllers.SendResponse(context, "Team has been created.")
	}
}

func AddMemberInTeam(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var teamMemberData models.TeamMember

		err1 := context.ShouldBindJSON(&teamMemberData)
		if err1 != nil {
			controllers.LogErrorWithAbort(context, err1, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err2 := database.Connect(server)
		if err2 != nil {
			controllers.LogErrorWithAbort(context, err2, http.StatusInternalServerError)
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			}
		}(db)

		// Check whether the member is a participant
		_, isMemberExists := GetUserDataByEmail(db, teamMemberData.MemberEmail)
		if !isMemberExists {
			controllers.LogErrorWithAbort(context, errors.New("no user found with this email"), http.StatusInternalServerError)
			return
		}

		_, err3 := db.Exec(database.InsertMemberInTeamQuery, teamMemberData.TeamId, teamMemberData.MemberEmail, teamMemberData.HackathonId)
		if err3 != nil {
			if strings.Contains(err3.Error(), "duplicate key") {
				controllers.LogErrorWithAbort(context, errors.New("the member is already in the team"), http.StatusInternalServerError)
				return
			}
			controllers.LogErrorWithAbort(context, err3, http.StatusInternalServerError)
			return
		}

		controllers.SendResponse(context, "Member has been added")
	}
}

func DeleteMemberFromTeam(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var memberInfo models.TeamMember

		err1 := context.ShouldBindJSON(&memberInfo)
		if err1 != nil {
			controllers.LogErrorWithAbort(context, err1, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err2 := database.Connect(server)
		if err2 != nil {
			controllers.LogErrorWithAbort(context, err2, http.StatusInternalServerError)
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
				return
			}
		}(db)

		// Delete participant data based on current logged in email
		_, err3 := db.Exec(database.DeleteMemberQuery, memberInfo.TeamId, memberInfo.MemberEmail)
		if err3 != nil {
			controllers.LogErrorWithAbort(context, err3, http.StatusInternalServerError)
			return
		}

		controllers.SendResponse(context, "Member has been delete successfully")
	}
}

// Get all team name you are a member
func GetAllTeamData(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
