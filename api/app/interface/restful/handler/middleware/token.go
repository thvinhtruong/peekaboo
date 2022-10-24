package middleware

import (
	"net/http"
	"server/app/interface/restful/handler/gctx"
	"server/utils/e"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ID         int
	EntityCode int
	jwt.StandardClaims
}

var jwtKey = []byte("peekaboo")

// About authentication mechanism:
// A JWT will be embeded inside the cookie. Thus, each JWT has expiration time of 1 week. After 1 week, user have to log-in again.
func GenerateJWT(c *gin.Context) string {
	expirationTime := time.Now().Add(24 * 7 * time.Hour)

	claims := &Claims{
		ID:         c.GetInt("ID"),
		EntityCode: c.GetInt("EntityCode"),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		app := gctx.Gin{C: c}

		cookies, err := app.C.Cookie("peekaboo")
		if err != nil {
			app.Response(http.StatusOK, 0, e.ErrorCookieNotFound)
			app.C.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(cookies, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			app.Response(http.StatusOK, 0, e.ErrorCookieNotFound)
			app.C.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("ID", claims.ID)
			c.Set("EntityCode", claims.EntityCode)
			c.Next()
			return

		} else {
			app.Response(http.StatusOK, 0, e.ErrorCookieOutdated)
			app.C.Abort()
			return
		}
	}
}
