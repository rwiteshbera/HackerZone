package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/HackZone/api"
	"net/http"
	"strings"
)

const (
	AuthorizationCookie     = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorizationPayload"
)

func AuthMiddleware(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken, err := context.Cookie(AuthorizationCookie)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization is not provided"})
			return
		}

		fields := strings.Fields(authToken)
		if len(fields) < 2 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization type", "type": authorizationType})
			return
		}

		accessToken := fields[1]
		payload, err := server.TokenMaker.VerifyToken(accessToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization token",
			})
		}

		context.Set(AuthorizationPayloadKey, payload)
		context.Next()
	}
}
