package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tot/core/schemes"
	"tot/tools/utils"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := utils.IsAuthorized(authToken, secret)
			if authorized {
				var userID, err = utils.ExtractToken(authToken, secret)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, schemes.ShmErrorResponse{
						Code: 97,
						Err:  "Not find User",
					})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, schemes.ShmErrorResponse{
				Code: 98,
				Err:  "Not authorized",
			})
			c.Abort()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, schemes.ShmErrorResponse{
			Code: 99,
			Err:  "Unable to find token in header",
		})
		c.Abort()
	}
}
