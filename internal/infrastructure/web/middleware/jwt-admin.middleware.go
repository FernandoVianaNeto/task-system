package middleware

import (
	"fmt"
	"net/http"
	"strings"
	configs "task-system/cmd/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtAdminClaims struct {
	Email string `json:"user_email"`
	Role  string `json:"user_role"`
	Uuid  string `json:"user_uuid"`
	jwt.RegisteredClaims
}

func JwtAdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := []byte(configs.ApplicationCfg.JwtSecret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token not provided"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &JwtAdminClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*JwtAdminClaims); ok && token.Valid {
			if claims.Role != "admin" {
				fmt.Println("CAI AQUI E O CARA N É ADMIN")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
				c.Abort()
				return
			}

			c.Set("user_email", claims.Email)
			c.Set("user_role", claims.Role)
			c.Set("user_uuid", claims.Uuid)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
		}
	}
}
