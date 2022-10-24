package middleware

import (
	"net/http"
	"server/app/interface/restful/handler/gctx"
	"server/utils/e"

	"github.com/gin-gonic/gin"
)

func AuthenticateUserRole(roles []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		app := gctx.Gin{C: c}

		ID := c.GetInt("ID")
		EntityCode := c.GetInt("EntityCode")

		if (ID == 0) || (EntityCode == 0) {
			app.Response(http.StatusUnauthorized, nil, e.ErrorNotAuthorized)
			app.C.Abort()
			return
		}

		for _, v := range roles {
			if v == EntityCode {
				c.Next()
				return
			}
		}

		app.Response(http.StatusUnauthorized, nil, e.ErrorNotAuthorized)
		app.C.Abort()
	}
}
