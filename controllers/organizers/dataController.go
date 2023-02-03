package organizers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/database"
	"github.com/rwiteshbera/HackZone/models"
	"net/http"
)

// GetOrganizerData : Get logged in Hackathon organizer's data
func GetOrganizerData(server *api.Server) gin.HandlerFunc {
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

		response, isExist := GetOrgData(db, currentEmail)
		if !isExist {
			logErrorWithAbort(context, errors.New("unable to find data"), http.StatusInternalServerError)
			return
		}

		var organizerDataResponse = &models.Participant{
			Email:     response.Email,
			FirstName: response.FirstName,
			LastName:  response.LastName,
			LastLogin: response.LastLogin,
			CreatedAt: response.CreatedAt,
		}

		context.JSON(http.StatusOK, gin.H{"message": organizerDataResponse})
	}
}
