package controllers

import (
	_ "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/database"
	"github.com/rwiteshbera/HackZone/models"
	"github.com/rwiteshbera/HackZone/utils"
	"net/http"
	"strings"
	"time"
)

func SignUp(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signUpRequest models.SignupUserRequest
		//var signUpResponse models.SignupUserResponse

		// Bind Request body into Object
		err := context.ShouldBindJSON(&signUpRequest)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusBadRequest)
			return
		}

		// Split full name into firstname and lastname
		name := strings.Split(signUpRequest.FullName, " ")

		hashedPassword, err1 := utils.HashPassword(signUpRequest.Password)
		if err1 != nil {
			logErrorWithAbort(context, err1, http.StatusInternalServerError)
			return
		}

		createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		lastLogin := createdAt

		// Connect with Database
		db, err := database.Connect(server)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusOK)
			return
		}

		tx := db.MustBegin()
		res, err := tx.Exec(database.SignUpQuery, signUpRequest.Email, name[0], name[1], hashedPassword, createdAt, lastLogin)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		err = tx.Commit()
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": res})
	}
}

func Login(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

// Handle error
func logErrorWithAbort(context *gin.Context, err error, statusCode int) {
	context.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}

// log a message
func logMessage(context *gin.Context, message any) {
	context.JSON(http.StatusOK, gin.H{"message": message})
	return
}
