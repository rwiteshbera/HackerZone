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
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				logErrorWithAbort(context, err, http.StatusInternalServerError)
			}
		}(db)

		// Check if user with the requested email already exists
		_, _, ifExist := CheckIfUserExists(db, signUpRequest.Email)
		if ifExist {
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
		var loginRequest models.LoginUserRequest
		//var loginResponse models.LoginUserResponse

		// Bind request body into JSON
		err := context.ShouldBindJSON(&loginRequest)
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

		// Check if user with the request email doesn't exist
		userEmail, hashedPassword, ifExists := CheckIfUserExists(db, loginRequest.Email)
		if !ifExists {
			logErrorWithAbort(context, errors.New("no user found"), http.StatusInternalServerError)
			return
		}

		// If user exists
		// Verify password
		isValid := utils.VerifyPassword(*hashedPassword, loginRequest.Password)
		if !isValid {
			logErrorWithAbort(context, errors.New("incorrect credentials"), http.StatusInternalServerError)
			return
		}

		tokenDuration, err := time.ParseDuration(server.Config.ACCESS_TOKEN_DURATION)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		accessToken, err := server.TokenMaker.CreateToken(*userEmail, tokenDuration)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		context.JSON(http.StatusOK, gin.H{"access_token": accessToken})

	}
}

// CheckIfUserExists : Check If user with this email already exists or not
func CheckIfUserExists(db *sql.DB, requestedEmail string) (*string, *string, bool) {
	var email, password string
	err := db.QueryRow(database.CheckIfUserAlreadyExistsQuery, requestedEmail).Scan(&email, &password)
	if err != nil {
		return nil, nil, false
	}
	if email == "" {
		return nil, nil, false
	}
	return &email, &password, true
}

// AddUserData : Add user data to database during signup
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
