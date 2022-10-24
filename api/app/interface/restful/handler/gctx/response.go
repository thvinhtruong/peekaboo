package gctx

import (
	"server/utils/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type ErrorResponse struct {
	ErrorCode  int    `json:"errorCode"`
	ErrorMsg   string `json:"errorMsg"`
	ErrorField string `json:"errorField"`
}

func classifyErrorResponse(GivenError error) (field string) {
	switch GivenError {
	case nil:
		return ""
	// The following are error field, not elaboration.
	case e.ErrorInvalidUsername:
		return "Username"
	case e.ErrorInvalidPassword:
		return "Password"
	case e.ErrorInvalidFullname:
		return "Fullname"
	case e.ErrorSpaceDetected:
		return "Username and Password"
	case e.ErrorCookieNotFound:
		return "Cookie"
	case e.ErrorBlockedIP:
		return "Blocked"
	case e.ErrorWrongDateFormat:
		return "Date"

		// The following are elaboration, not error field.
	case e.ErrorEntryExists:
		return "The object you are trying to create already existed. This can be fix when you change the Username so that it unique"
	case e.ErrorCookieOutdated:
		return "Your session expired. Please login again"
	case e.ErrorInputInvalid:
		return "Please notice your input to be valid"
	case e.ErrorNotAuthorized:
		return "You are not authorized to do this action"
	case e.ErrorNotAuthenticated:
		return "You are not authenticated, please login"
	case e.ErrorInternal:
		return "The system is suffering an issue, please contact the admin with record of behaviour"
	case e.ErrorEntryNotExist:
		return "The system is suferring an inconsistency in behavior, please contact the admin with record of behaviour"
	case e.ErrorBindJSON:
		return "The system connection is unstable, please contact the admin with record of behaviour"
	case e.ErrorContextLostValue:
		return "The system authentication of your account has been deleted, please login again"
	case e.ErrorNameAlreadyExist:
		return "choose another name for your object (new book name, new class name, new username, etc)"
	default:
		return "Unidentified Error, please contact the admin with record of behaviour"
	}
}

func (g *Gin) Response(httpCode int, data interface{}, GivenError error) {
	var msg string
	if GivenError == nil {
		msg = ""
	} else {
		msg = GivenError.Error()
	}

	field := classifyErrorResponse(GivenError)
	code := e.GetCode(GivenError)

	g.C.Header("Content-Type", "application/json")
	g.C.JSON(httpCode, gin.H{
		"data": data,
		"error": ErrorResponse{
			ErrorCode:  code,
			ErrorMsg:   msg,
			ErrorField: field,
		},
	})
}
