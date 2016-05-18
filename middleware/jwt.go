package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func JWTVerify(key string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // parse jwt
        token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
            return []byte(key), nil
        })

        if err == nil && token.Valid {
            // move on
            c.Next()
        } else {
            // abort
            c.AbortWithStatus(http.StatusUnauthorized)
        }
    }
}
