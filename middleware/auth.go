package middleware
//this file job is the before allowing a user to access protected routers, check whether the jwt token is valid or not
import (
	"net/http"
	"simple-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc { //this measn return a function that handles request.
	return func(c *gin.Context) { //this this is the actual middleware every request passes through this.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorizationheadermissing"})
			c.Abort() //do not go the next handler.
			return
		}

		//excecting formate: Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) !=2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invaid authorization header formate"})
			c.Abort()
			return
		}

		tokenStr := parts[1] //take only the token part, tokenStr contines a 
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		//pass user info into context for header to use
		c.Set("user_id", claims.UserID) //save data in gin context temporarly for the next handler.
		c.Set("email", claims.Email)

		c.Next()
	}
}