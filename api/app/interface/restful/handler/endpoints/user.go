package endpoints

import (
	"context"
	"log"
	"net/http"
	"server/app/interface/persistence/rdbms/sqlconnection"
	"server/app/interface/restful/handler/api_dto"
	"server/app/interface/restful/handler/gctx"
	"server/app/interface/restful/handler/middleware"
	"server/app/registry"
	"server/app/usecase/usecase_dto"
	"server/utils/e"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func UserHandler(c *gin.RouterGroup) {
	c.GET("/", GetUserInfo)
	c.GET("/all", GetAllUsers)
	c.GET("/class", GetUserClass)
	c.GET("/test_answer", ReviewTestAnswer)
	c.GET("/test_result", GetAllTestResult)

	c.PUT("/", UpdateUser)
}

func GetUserInfo(c *gin.Context) {
	var user_record usecase_dto.User
	app := gctx.Gin{C: c}
	ctx := context.Background()

	username := c.Query("username")
	fullname := c.Query("fullname")
	userId := c.Query("user_id")
	if len(username) == 0 && len(fullname) == 0 && len(userId) == 0 {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	var ID int
	var err error

	if len(userId) != 0 {
		ID, err = strconv.Atoi(userId)
		if err != nil {
			app.Response(http.StatusInternalServerError, 0, err)
			return
		}
	}

	if username != "" || fullname != "" || ID != 0 {
		user_record = usecase_dto.User{ID: ID, Username: username, FullName: fullname}
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	user, err := access.Service.FindUser(ctx, user_record, false)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, user, nil)

}

func GetAllUsers(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	user, err := access.Service.FindAllUser(ctx)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var users []api_dto.User
	if err := copier.Copy(&users, &user); err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, users, nil)
}

func GetUserClass(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	userID := c.Query("user_id")
	if userID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}
	ID, err := strconv.Atoi(userID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	classes, err := access.Service.FindUserClasses(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, classes, nil)
}

func ReviewTestAnswer(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	resultID := c.Query("result_id")
	if resultID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}
	ID, err := strconv.Atoi(resultID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	result, err := access.Service.ReviewTestResult(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, result, nil)
}

func GetAllTestResult(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	userId := c.Query("user_id")
	if userId == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	ID, err := strconv.Atoi(userId)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	result, err := access.Service.FindAllUserTestResult(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, result, nil)
}

func UpdateUser(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	data, err := api_dto.BindUserModel(c, true)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var user_record usecase_dto.User
	err = copier.Copy(&user_record, data)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	err = access.Service.UpdateUser(ctx, user_record)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, 1, nil)
}

func CreateUser(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	data, err := api_dto.BindUserModel(c, false)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	log.Println("data: ", data)

	var user_record usecase_dto.User
	err = copier.Copy(&user_record, &data)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	password, err := middleware.HashPassword(user_record.Password)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	user_record.Password = password
	user_record.EntityCode = 3

	log.Println("record: ", user_record)

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	IDs, err := access.Service.CreateUser(ctx, user_record)
	if err != nil {
		log.Println("err: ", err)
		app.Response(http.StatusOK, 0, err)
		return
	}

	app.Response(http.StatusOK, IDs, nil)
}
