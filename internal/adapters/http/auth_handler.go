package http

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var rsaPublicKey *rsa.PublicKey

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("WARN: No .env file, using default system variables")
	}
	publicKeyPath := os.Getenv("RSA_PUBLIC_KEY")
	if publicKeyPath == "" {
		log.Fatal("FATAL: RSA_PUBLIC_KEY environment variable not set.")
	}

	keyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("FATAL: could not read public key file at path '%s': %v", publicKeyPath, err)
	}

	rsaPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		log.Fatalf("FATAL: could not parse PEM-encoded public key: %v", err)
	}
	log.Println("Successfully loaded RSA public key for JWT verification.")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token format is required"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return rsaPublicKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if email, ok := claims["email"].(string); ok {
				c.Set("email", email)
			}

			if sub, ok := claims["sub"].(string); ok {
				num, err := strconv.Atoi(sub)

				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
					return
				}

				c.Set("userId", num)
				c.Next()

				return
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "'email' claim not found or is not a string"})
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
	}
}
