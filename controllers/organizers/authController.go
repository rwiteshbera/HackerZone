package organizers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/HackZone/api"
	"github.com/rwiteshbera/HackZone/database"
	"github.com/rwiteshbera/HackZone/middlewares"
	"github.com/rwiteshbera/HackZone/models"
	"github.com/rwiteshbera/HackZone/utils"
	"net/http"
	"strings"
	"time"
)

// SignUpAsOrganizer : Join as hackathon organizer
func SignUpAsOrganizer(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signUpRequest models.SignupRequest
		//var signUpResponse models.SignupUserResponse

		// Bind Request body into Object
		err1 := context.ShouldBindJSON(&signUpRequest)
		if err1 != nil {
			logErrorWithAbort(context, err1, http.StatusBadRequest)
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

		// Check if organizer with the requested email already exists
		_, isExist := GetOrgData(db, signUpRequest.Email)
		if isExist {
			logErrorWithAbort(context, errors.New("a person with the same email has already established an organizer account"), http.StatusInternalServerError)
			return
		}

		// If user doesn't exist, then add user data
		err3 := AddOrganizerData(db, signUpRequest)
		if err3 != nil {
			logErrorWithAbort(context, err3, http.StatusInternalServerError)
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "registration successful"})

	}
}

// LoginAsOrganizer : Login as a hackathon organizer
func LoginAsOrganizer(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginRequest models.LoginRequest

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
		hostDataResponse, isExist := GetOrgData(db, loginRequest.Email)
		if !isExist {
			logErrorWithAbort(context, errors.New("no user found"), http.StatusInternalServerError)
			return
		}

		// If user exists
		// Verify password
		isValid := utils.VerifyPassword(hostDataResponse.Password, loginRequest.Password)
		if !isValid {
			logErrorWithAbort(context, errors.New("incorrect credentials"), http.StatusInternalServerError)
			return
		}

		tokenDuration, err := time.ParseDuration(server.Config.ACCESS_TOKEN_DURATION)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		accessToken, err := server.TokenMaker.CreateToken(hostDataResponse.Email, tokenDuration)
		if err != nil {
			logErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		var loginResponse = models.LoginResponse{
			Email:     hostDataResponse.Email,
			FirstName: hostDataResponse.FirstName,
			LastName:  hostDataResponse.LastName,
			LastLogin: hostDataResponse.LastLogin,
			CreatedAt: hostDataResponse.CreatedAt}

		context.SetCookie("authorization", middlewares.AuthorizationTypeBearer+" "+accessToken, 0, "/", server.Config.SERVER_HOST, false, true)
		context.SetCookie("email", loginResponse.Email, 0, "/", server.Config.SERVER_HOST, false, true)
		context.JSON(http.StatusOK, gin.H{"message": loginResponse})
	}
}

// GetOrgData : Check If organizer with this email already exists or not, if exists return all the data
// return firstname, lastname, email, password, lastLogin, CreatedAt
func GetOrgData(db *sql.DB, requestedEmail string) (*models.Participant, bool) {
	var response models.Participant
	err := db.QueryRow(database.GetOrganizersDataQuery, requestedEmail).Scan(&response.Email, &response.FirstName, &response.LastName, &response.Password, &response.LastLogin, &response.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, false
	}
	return &response, true
}

// AddOrganizerData : Add user data to database during signup
func AddOrganizerData(db *sql.DB, request models.SignupRequest) error {

	// Split full name into firstname and lastname
	name := strings.Split(request.FullName, " ")

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	lastLogin := createdAt

	_, err = db.Exec(database.SignupAsOrganizerQuery, request.Email, name[0], name[1], hashedPassword, lastLogin, createdAt)
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