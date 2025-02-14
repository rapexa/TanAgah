// internal/middlewares/jwt_middleware.go
package middlewares

import (
	"TanAgah/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTMiddleware is a middleware function for validating JWT.
func JWTMiddleware(c *gin.Context) {
	// Extract token from the Authorization header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort() // Stop the request
		return
	}

	// Validate the token
	ok, err := utils.ValidateToken(token)
	if err != nil && !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort() // Stop the request
		return
	}

	c.Next() // Proceed to the next handler
}
