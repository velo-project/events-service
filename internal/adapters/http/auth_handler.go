package http

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
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

func AuthMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Cabeçalho de autorização é obrigatório", "status_code": http.StatusUnauthorized})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Formato de token Bearer é obrigatório", "status_code": http.StatusUnauthorized})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return rsaPublicKey, nil
		})

		if err != nil {
			sentry.CaptureException(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token inválido: " + err.Error(), "status_code": http.StatusUnauthorized})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if len(roles) > 0 {
				if role, ok := claims["role"].(string); ok {
					hasRole := false
					for _, r := range roles {
						if r == role {
							hasRole = true
							break
						}
					}
					if !hasRole {
						c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Proibido", "status_code": http.StatusForbidden})
						return
					}
				} else {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Proibido", "status_code": http.StatusForbidden})
					return
				}
			}

			if email, ok := claims["email"].(string); ok {

				c.Set("email", email)
				if hub := sentrygin.GetHubFromContext(c); hub != nil {
					hub.WithScope(func(scope *sentry.Scope) {
						scope.SetExtra("email", email)
					})
				}
			}

			if sub, ok := claims["sub"].(string); ok {
				num, err := strconv.Atoi(sub)

				if err != nil {
					sentry.CaptureException(err)
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token inválido: " + err.Error(), "status_code": http.StatusUnauthorized})
					return
				}

				c.Set("userId", num)
				if hub := sentrygin.GetHubFromContext(c); hub != nil {
					hub.WithScope(func(scope *sentry.Scope) {
						scope.SetExtra("user_id", num)
					})
				}
				c.Next()

				return
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Claim 'email' não encontrada ou não é uma string", "status_code": http.StatusUnauthorized})
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Claims de token inválidas", "status_code": http.StatusUnauthorized})
	}
}
