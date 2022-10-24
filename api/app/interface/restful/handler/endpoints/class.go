package endpoints

import (
	"context"
	"net/http"
	"server/app/interface/persistence/rdbms/sqlconnection"
	"server/app/interface/restful/handler/api_dto"
	"server/app/interface/restful/handler/gctx"
	"server/app/registry"
	"server/app/usecase/usecase_dto"
	"server/utils/e"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func ClassHandler(c *gin.RouterGroup) {
	c.GET("/", GetAllClasses)
	c.GET("/tests", GetClassTest)
	c.GET("/single_test", GetSingleClassTest)
	c.GET("/members", GetClassMembers)

	c.POST("/", CreateClass)
	c.POST("/add_member", AddMember)
	c.POST("/add_test", AddTest)

	c.DELETE("/remove_test", RemoveTest)
	c.DELETE("/remove_member", RemoveMember)
}

func GetSingleClassTest(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	CID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	testID := c.Query("test_id")
	if testID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	TID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	test, err := access.Service.GetSingleClassTest(ctx, CID, TID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var result api_dto.Test
	if err := copier.Copy(&result, &test); err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, result, nil)
}

func GetAllClasses(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	classes, err := access.Service.GetClasses(ctx)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, classes, nil)
}

func GetClassTest(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	ID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	tests, err := access.Service.GetClassTest(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var test []api_dto.Test
	if err := copier.Copy(&test, &tests); err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, test, nil)
}

func CreateClass(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	data, err := api_dto.BindClass(c, false, true)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var class usecase_dto.Class
	if err := copier.Copy(&class, &data); err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	id, err := access.Service.CreateClass(ctx, class)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, id, nil)
}

func GetClassMembers(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	ID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	members, err := access.Service.QueryClassMembers(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, members, nil)
}

func AddMember(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	CID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	UID, err := strconv.Atoi(userID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	err = access.Service.AddMember2Class(ctx, CID, UID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, 1, nil)
}

func RemoveMember(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	CID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	UID, err := strconv.Atoi(userID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	err = access.Service.RemoveMemberFromClass(ctx, CID, UID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, 1, nil)
}

func AddTest(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	CID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	testId := c.Query("test_id")
	if testId == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	TID, err := strconv.Atoi(testId)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	err = access.Service.AddTest2Class(ctx, CID, TID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, 1, nil)
}

func RemoveTest(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	classID := c.Query("class_id")
	if classID == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	CID, err := strconv.Atoi(classID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	testId := c.Query("test_id")
	if testId == "" {
		app.Response(http.StatusOK, 0, e.ErrorInputInvalid)
		return
	}

	TID, err := strconv.Atoi(testId)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildClassAccessPoint(false, sqlconnection.DBConn)
	err = access.Service.RemoveTestClass(ctx, CID, TID)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, 1, nil)
}
