package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwksURL     = "https://keycloak.identialab.com/realms/pqr/protocol/openid-connect/certs"
	expectedIss = "https://keycloak.identialab.com/realms/pqr"
	expectedAud = "backend-api"
	jwks        *keyfunc.JWKS
)

func init() {
	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Hour,
	})
	if err != nil {
		panic("No se pudo obtener las llaves JWKS: " + err.Error())
	}
}

// AuthMiddleware es compatible con Gin
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó el token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, jwks.Keyfunc)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		// Validar issuer
		if iss, ok := claims["iss"].(string); !ok || iss != expectedIss {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Issuer inválido"})
			return
		}

		// Validar audience
		aud, ok := claims["aud"].(string)
		if !ok || (aud != "frontend" && aud != "account") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Audience inválido"})
			return
		}

		// Guardar claims en el contexto para usarlos luego
		c.Set("claims", claims)

		c.Next()
	}
}
