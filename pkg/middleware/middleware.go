package middleware

import (
	"TanAgah/internal/repository" // Import the repository package
	"TanAgah/internal/utils"      // Adjust the import path as needed
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTMiddleware is a middleware function for validating JWT.
func JWTMiddleware(repo repository.JWTRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort() // Stop the request
			return
		}

		// Validate the token
		ok, err := utils.ValidateToken(token)
		if err != nil || !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // Stop the request
			return
		}

		// Check the token in the database
		if !repo.IsTokenValid(token) { // This function should be implemented in your repository
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found or invalid"})
			c.Abort() // Stop the request
			return
		}

		c.Next() // Proceed to the next handler
	}
}
