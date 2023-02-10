package organizers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/database"
	"github.com/rwiteshbera/HackerZone/models"
	"net/http"
)

// Create/ Host a new hackathon as an organizer
func CreateHack(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var hackathonInfo models.Hackathon

		createdBy, err := context.Cookie("user")
		if err != nil {
			logErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}
		hackathonInfo.CreatedBy = createdBy

		err = context.ShouldBindJSON(&hackathonInfo)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		// Connect with Database
		db, err2 := database.Connect(server)
		if err2 != nil {
			logErrorWithAbort(context, err2, http.StatusInternalServerError)
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				logErrorWithAbort(context, err, http.StatusInternalServerError)
			}
		}(db)

		var hackNameFromDB string
		err = db.QueryRow(database.HackathonNameCheckQuery, hackathonInfo.Name).Scan(&hackNameFromDB)
		if err != nil {
			if err == sql.ErrNoRows {
				// Hackathon name is available
				// Create new
				_, err = db.Exec(database.CreateHackathonQuery, hackathonInfo.Name, hackathonInfo.Tagline, hackathonInfo.Description, hackathonInfo.ContactEmail, hackathonInfo.Host, hackathonInfo.HackingStart, hackathonInfo.Deadline, hackathonInfo.CreatedBy)
				return
			}
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"message": "hackathon name is already used!"})
	}
}
