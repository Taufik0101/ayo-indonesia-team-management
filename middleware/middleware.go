package middleware

import (
	"gin-ayo/config"
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		accessTokenSecret := config.GetEnv("JWT_SECRET", "")
		validate, err := utils.JwtValidate(tokenString, accessTokenSecret)
		if err != nil || !validate.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		extractedClaims, err := utils.ExtractClaims(*validate)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		var accessTokenPayload dto.AccessTokenPayload
		err = mapstructure.Decode(extractedClaims["payload"], &accessTokenPayload)
		if err != nil {
			log.Println("Error decoding payload:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Payload Malformed"})
			c.Abort()
			return
		}

		c.Set("user", &accessTokenPayload)
		c.Next()
	}
}
