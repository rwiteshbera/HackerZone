package participants

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/database"
	"github.com/rwiteshbera/HackerZone/models"
	"net/http"
)

// GetParticipantData : Get logged in participants data
func GetParticipantData(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		currentEmail, err := context.Cookie("email")
		if err != nil {
			logErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err := database.Connect(server)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				logErrorWithAbort(context, err, http.StatusInternalServerError)
				return
			}
		}(db)

		response, isExist := GetUserData(db, currentEmail)
		if !isExist {
			logErrorWithAbort(context, errors.New("unable to find participant"), http.StatusInternalServerError)
			return
		}

		var useDataResponse = &models.User{
			UUID:      response.UUID,
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
			logErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err := database.Connect(server)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				logErrorWithAbort(context, err, http.StatusInternalServerError)
				return
			}
		}(db)

		// Delete participant data based on current logged in email
		_, err = db.Exec(database.DeleteParticipantQuery, currentEmail)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "You profile has been delete successfully"})
	}
}
