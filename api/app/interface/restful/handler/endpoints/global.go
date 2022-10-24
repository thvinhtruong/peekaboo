package endpoints

import (
	"context"
	"net/http"
	"server/app/interface/persistence/rdbms/sqlconnection"
	"server/app/interface/restful/handler/api_dto"
	"server/app/interface/restful/handler/gctx"
	"server/app/interface/restful/handler/middleware"
	"server/app/registry"
	"server/app/usecase/usecase_dto"
	"server/setting"
	"server/utils/e"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Login(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()
	UserAC := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)

	data, err := api_dto.BindUserModel(c, true)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var user usecase_dto.User
	if err := copier.Copy(&user, &data); err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
	}

	records, err := UserAC.Service.FindUser(ctx, user, true)
	if err != nil {
		app.Response(http.StatusOK, 0, err)
		return
	}

	if len(records) == 0 {
		app.Response(http.StatusOK, 0, e.ErrorEntryNotExist)
		return
	}

	if err := middleware.ComparePassword(records[0].Password, user.Password); err != nil {
		app.Response(http.StatusOK, 0, e.ErrorPasswordIncorrect)
		return
	}

	records[0].Password = ""

	c.Set("ID", records[0].ID)
	c.Set("EntityCode", records[0].EntityCode)

	jwt := middleware.GenerateJWT(c)
	// TODO: set https for testing.
	app.C.SetCookie("peekaboo", jwt, 604800, "/", setting.CookieDomain, setting.CookieSecure, setting.CookieHTTPS)

	var userModel api_dto.User
	if err := copier.Copy(&userModel, &records[0]); err != nil {
		app.Response(http.StatusOK, 0, err)
		return
	}

	app.Response(http.StatusOK, userModel, nil)
}

func Logout(c *gin.Context) {
	app := gctx.Gin{C: c}
	app.C.SetCookie("peekaboo", "", -1, "/", setting.CookieDomain, setting.CookieSecure, setting.CookieHTTPS)
	app.Response(http.StatusOK, 1, nil)
}

func ValidateRole(c *gin.Context) {
	app := gctx.Gin{C: c}
	entityCode := c.GetInt("EntityCode")

	app.Response(http.StatusOK, entityCode, nil)
}

func GetID(c *gin.Context) {
	app := gctx.Gin{C: c}
	ID := c.GetInt("ID")

	app.Response(http.StatusOK, ID, nil)
}
