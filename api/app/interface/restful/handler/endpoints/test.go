package endpoints

import (
	"context"
	"errors"
	"log"
	"net/http"
	"server/app/interface/persistence/rdbms/sqlconnection"
	"server/app/interface/restful/handler/api_dto"
	"server/app/interface/restful/handler/gctx"
	"server/app/registry"
	"server/app/usecase/usecase_dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func TestHandler(c *gin.RouterGroup) {
	c.GET("/", GetTestInfo)
	c.GET("/all", GetAllTests)
	c.GET("/result", GetTestResult)
	c.GET("/answer", GetTestAnswer)
	c.GET("/do", DoTest)

	c.POST("/submit", SubmitTest)
}

// Returns the title,
func GetTestInfo(c *gin.Context) {
	ctx := context.Background()
	app := gctx.Gin{C: c}

	testID := c.Query("test_id")
	if len(testID) == 0 {
		app.Response(http.StatusInternalServerError, nil, errors.New("test_id is required, but empty."))
		return
	}

	ID, err := strconv.Atoi(testID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	UserID := c.GetInt("ID")

	access := registry.BuildTestAccessPoint(false, sqlconnection.DBConn)
	test, err := access.Service.QueryTestInfo(ctx, ID, UserID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	var apitest api_dto.Test
	if err := copier.Copy(&apitest, test); err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	app.Response(http.StatusOK, apitest, nil)
}

func GetAllTests(c *gin.Context) {
	app := gctx.Gin{C: c}
	ctx := context.Background()

	access := registry.BuildTestAccessPoint(false, sqlconnection.DBConn)
	test, err := access.Service.QueryAllTest(ctx)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	var tests []api_dto.Test
	if err := copier.Copy(&tests, &test); err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	app.Response(http.StatusOK, tests, nil)
}

func GetTestResult(c *gin.Context) {
	ctx := context.Background()
	app := gctx.Gin{C: c}

	testResultID := c.Query("test_result_id")
	if len(testResultID) == 0 {
		app.Response(http.StatusInternalServerError, nil, errors.New("test_result_id is required, but empty."))
		return
	}

	ID, err := strconv.Atoi(testResultID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	access := registry.BuildTestResultAccessPoint(false, sqlconnection.DBConn)
	test, err := access.Service.GetTestResultHeadline(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	app.Response(http.StatusOK, test, nil)
}

func GetTestAnswer(c *gin.Context) {
	ctx := context.Background()
	app := gctx.Gin{C: c}

	testResultID := c.Query("test_result_id")
	if len(testResultID) == 0 {
		app.Response(http.StatusInternalServerError, nil, errors.New("test_result_id is required, but empty."))
		return
	}

	ID, err := strconv.Atoi(testResultID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	access := registry.BuildTestAccessPoint(false, sqlconnection.DBConn)
	test, err := access.Service.QueryTestAnswer(ctx, ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	app.Response(http.StatusOK, test, nil)
}

func DoTest(c *gin.Context) {
	app := gctx.Gin{C: c}

	testId := c.Query("test_id")
	if len(testId) == 0 {
		app.Response(http.StatusInternalServerError, nil, errors.New("test_id is required, but empty."))
		return
	}

	ID, err := strconv.Atoi(testId)
	if err != nil {
		app.Response(http.StatusInternalServerError, 0, err)
		return
	}

	access := registry.BuildTestAccessPoint(false, sqlconnection.DBConn)
	test, err := access.Service.QuerySkillTest(context.Background(), ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	var apitest api_dto.SkillTest
	if err := copier.Copy(&apitest, &test); err != nil {
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	app.Response(http.StatusOK, apitest, nil)
}

func SubmitTest(c *gin.Context) {
	ctx := context.Background()
	app := gctx.Gin{C: c}

	data, err := api_dto.BindSubmitData(c)
	if err != nil {
		log.Println(err)
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	var submitTest usecase_dto.SubmitData
	if err := copier.Copy(&submitTest, &data); err != nil {
		log.Println(err)
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	ID := c.GetInt("ID")
	EntityCode := c.GetInt("ID")

	access := registry.BuildTestAccessPoint(true, sqlconnection.DBConn)
	resultID, err := access.Service.SubmitTest(ctx, submitTest, ID, EntityCode)
	if err != nil {
		log.Println(err)
		app.Response(http.StatusInternalServerError, nil, err)
		return
	}

	app.Response(http.StatusOK, resultID, nil)
}
