package middleware

import (
	"errors"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// identity holds the verified claims we expose to handlers.
type identity struct {
	userID int64
	email  string
	name   string
}

// parseIdentity validates the token (HS256 only, exp required) and extracts
// the identity claims. It never panics on missing or mistyped claims.
func parseIdentity(tokenStr, jwtSecret string) (*identity, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}), jwt.WithExpirationRequired())
	if err != nil || token == nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	uid, ok := claims["user_id"].(float64)
	if !ok || uid < 1 || uid != math.Trunc(uid) || uid > math.MaxInt64/2 {
		return nil, errors.New("invalid user_id claim")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email claim")
	}
	name, ok := claims["name"].(string)
	if !ok {
		return nil, errors.New("invalid name claim")
	}
	return &identity{userID: int64(uid), email: email, name: name}, nil
}

func (id *identity) apply(c *gin.Context) {
	c.Set("user_id", id.userID)
	c.Set("email", id.email)
	c.Set("name", id.name)
}

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		id, err := parseIdentity(strings.TrimPrefix(authHeader, "Bearer "), jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		id.apply(c)
		c.Next()
	}
}

func AuthOptional(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}
		if id, err := parseIdentity(strings.TrimPrefix(authHeader, "Bearer "), jwtSecret); err == nil {
			id.apply(c)
		}
		c.Next()
	}
}
