package server

import (
	"net/http"
	"savegen-api/util"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "auth_user_id"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			respondUnauthorized(c, "Missing authorization header")
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			respondUnauthorized(c, "Invalid authorization header")
			return
		}

		claims, err := util.ParseToken(parts[1])
		if err != nil {
			respondUnauthorized(c, "Invalid or expired token")
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Next()
	}
}

func respondUnauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    "UNAUTHORIZED",
		"message": message,
		"data":    nil,
	})
}

func MustUserID(c *gin.Context) (int, bool) {
	raw, ok := c.Get(ContextUserIDKey)
	if !ok {
		return 0, false
	}
	id, ok := raw.(int)
	return id, ok
}
