package middleware

import (
	"TanAgah/internal/repository" // Import the repository package
	"TanAgah/internal/stringResource"
	"TanAgah/internal/utils" // Adjust the import path as needed
	"github.com/gin-gonic/gin"
)

// JWTMiddleware is a middleware function for validating JWT.
func JWTMiddleware(repo repository.JWTRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.SendError401Response(c, stringResource.GetStrings().TokenJwtIsRequired(c))
			c.Abort() // Stop the request
			return
		}

		// Validate the token
		ok, err := utils.ValidateToken(token)
		if err != nil || !ok {
			utils.SendError401Response(c, stringResource.GetStrings().TokenJwtIsNotValid(c))
			c.Abort() // Stop the request
			return
		}

		// Check the token in the database
		if !repo.IsTokenValid(token) { // This function should be implemented in your repository
			utils.SendError401Response(c, stringResource.GetStrings().UserNotFound(c))
			c.Abort() // Stop the request
			return
		}

		c.Next() // Proceed to the next handler
	}
}
