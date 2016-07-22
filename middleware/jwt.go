package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// JWTVerify middleware verifys a key and moves on if correct
func JWTVerify(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse jwt
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if err == nil && token.Valid {
			// set claims to access
			c.Set("claims", token.Claims)

			// move on
			c.Next()
		} else {
			// abort
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
