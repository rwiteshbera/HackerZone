package participants

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rwiteshbera/HackerZone/api"
	"github.com/rwiteshbera/HackerZone/controllers"
	"github.com/rwiteshbera/HackerZone/database"
	"github.com/rwiteshbera/HackerZone/middlewares"
	"github.com/rwiteshbera/HackerZone/models"
	"github.com/rwiteshbera/HackerZone/utils"
	"net/http"
	"strings"
	"time"
)

func SignUp(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signUpRequest models.SignupRequest

		// Bind Request body into Object
		err1 := context.ShouldBindJSON(&signUpRequest)
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

		// Check if user with the requested email already exists
		_, isExist := GetUserDataByEmail(db, signUpRequest.Email)
		if isExist {
			controllers.LogErrorWithAbort(context, errors.New("user already exists"), http.StatusInternalServerError)
			return
		}

		// If user doesn't exist, then add user data
		err3 := AddUserData(db, signUpRequest)
		if err3 != nil {
			controllers.LogErrorWithAbort(context, err3, http.StatusInternalServerError)
			return
		}

		controllers.SendResponse(context, "registration successful")
	}
}

func Login(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginRequest models.LoginRequest

		// Bind request body into JSON
		err := context.ShouldBindJSON(&loginRequest)
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

		// Check if user with the request email doesn't exist
		userDataResponse, isExist := GetUserDataByEmail(db, loginRequest.Email)
		if !isExist {
			controllers.LogErrorWithAbort(context, errors.New("no user found"), http.StatusInternalServerError)
			return
		}

		// If user exists
		// Verify password
		isValid := utils.VerifyPassword(userDataResponse.Password, loginRequest.Password)
		if !isValid {
			controllers.LogErrorWithAbort(context, errors.New("incorrect credentials"), http.StatusInternalServerError)
			return
		}

		tokenDuration, err := time.ParseDuration(server.Config.ACCESS_TOKEN_DURATION)
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		accessToken, err := server.TokenMaker.CreateToken(userDataResponse.Email, tokenDuration)
		if err != nil {
			controllers.LogErrorWithAbort(context, err, http.StatusInternalServerError)
			return
		}

		var loginResponse = models.LoginResponse{
			FirstName: userDataResponse.FirstName,
			LastName:  userDataResponse.LastName,
			Email:     userDataResponse.Email,
			Bio:       userDataResponse.Bio,
			Gender:    userDataResponse.Gender,
			LastLogin: userDataResponse.LastLogin,
			CreatedAt: userDataResponse.CreatedAt}

		context.SetCookie("authorization", middlewares.AuthorizationTypeBearer+" "+accessToken, 0, "/", server.Config.SERVER_HOST, false, true)
		context.SetCookie("email", loginResponse.Email, 0, "/", server.Config.SERVER_HOST, false, true)
		context.SetCookie("user", userDataResponse.UUID.String(), 0, "/", server.Config.SERVER_HOST, false, true)

		controllers.SendResponse(context, loginResponse)

	}
}

// GetUserData : Check If user with this email already exists or not, if exists return all the data
// return firstname, lastname, email, password, lastLogin, CreatedAt
func GetUserDataByEmail(db *sql.DB, reqEmail string) (*models.User, bool) {
	var response models.User
	err := db.QueryRow(database.GetUserAllDataByEmailQuery, reqEmail).Scan(&response.UUID, &response.Email, &response.FirstName, &response.LastName, &response.Bio, &response.Gender, &response.Password, &response.LastLogin, &response.CreatedAt)
	if err != nil {
		return nil, false
	}
	if err == sql.ErrNoRows {
		return nil, false
	}
	return &response, true
}

// AddUserData : Add user data to database during signup
func AddUserData(db *sql.DB, request models.SignupRequest) error {

	// Split full name into firstname and lastname
	name := strings.Split(request.FullName, " ")

	// Hash the password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	// Create UUID
	var uuid_new = uuid.New()

	// Add time
	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	lastLogin := createdAt

	_, err = db.Exec(database.SignUpQuery, uuid_new, request.Email, name[0], name[1], request.Bio, request.Gender, hashedPassword, lastLogin, createdAt)
	if err != nil {
		return err
	}
	return nil
}
