package controllers

import (
	"database/sql"
	_ "database/sql"
	"errors"
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

		// Connect with Database
		db, err := database.Connect(server)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Check if user with the requested email already exists
		if CheckIfUserExists(db, signUpRequest) {
			logErrorWithAbort(context, errors.New("user already exists"), http.StatusInternalServerError)
			return
		}

		// If email doesn't exist, then add user data
		err = AddUserData(db, signUpRequest)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "registration successful"})

	}
}

func Login(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

// Check If user with this email already exists or not
func CheckIfUserExists(db *sql.DB, request models.SignupUserRequest) bool {
	var email string
	db.QueryRow(database.CheckIfUserAlreadyExistsQuery, request.Email).Scan(&email)
	if email == "" {
		return false
	}
	return true
}

// Add user data to database during signup
func AddUserData(db *sql.DB, request models.SignupUserRequest) error {

	// Split full name into firstname and lastname
	name := strings.Split(request.FullName, " ")

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	lastLogin := createdAt

	_, err = db.Exec(database.SignUpQuery, request.Email, name[0], name[1], hashedPassword, createdAt, lastLogin)
	if err != nil {
		return err
	}
	return nil
}

// Handle error
func logErrorWithAbort(context *gin.Context, err error, statusCode int) {
	context.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error(), "package": "controllers"})
}

// log a message
func logMessage(context *gin.Context, message any) {
	context.JSON(http.StatusOK, gin.H{"message": message})
}
