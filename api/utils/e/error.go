package e

import "errors"

// Validation Error
// Code start with 1xxxxx
var (
	ErrorInvalidUsername = errors.New("invalid Username, please try again")

	ErrorInvalidPassword = errors.New("invalid Password, please try again")

	ErrorInvalidFullname = errors.New("invalid Fullname, please try again")

	ErrorInvalidPhoneNumber = errors.New("invalid Phone number, please try again")

	ErrorEmptyField = errors.New("empty field")

	ErrorSpaceDetected = errors.New("there should be no space on Username / Password field")

	ErrorWrongDateFormat = errors.New("wrong given format for date")

	ErrorPasswordIncorrect = errors.New("password is incorrect")
)

// Client errors.
// Code start with 2xxxx
var (
	// use when try to create but fail as entry existed.
	ErrorEntryExists = errors.New("record already exists")

	// wrong input.
	ErrorInputInvalid = errors.New("input invalid")

	// use when cookie is not found.
	ErrorCookieNotFound = errors.New("cookie not found")

	// use when response to a blocked IP address.
	ErrorBlockedIP = errors.New("blocked IP")

	// use when authentication failed
	ErrorNotAuthenticated = errors.New("authentication failed")

	ErrorNotAuthorized = errors.New("not authorized for this action ")
	// success

	ErrorCookieOutdated = errors.New("cookie outdated")

	ErrorCannotDeleteThisEntity = errors.New("cannot delete this user because it will violate the system")

	ErrorRelationViolation = errors.New("cannot add this relationship the current user, please check that you select the right users and/or objects of this action")

	ErrorCannotDeleteGuardian = errors.New("student must have at least 1 guardians")

	ErrorUsernameAlreadyExist = errors.New("username already exists, please choose another username")

	ErrorClassnameAlreadyExist = errors.New("classname already exists, please choose another name")

	ErrorNameAlreadyExist = errors.New("the name of the object has already existed. Please choose another name")
)

// Server errors.
var (
	// use when error is not related to a user but rather server error.
	ErrorInternal = errors.New("error internal")

	// use whe try to query/process query but fail as some kinds of relations is wrong.
	// need to taken care by db master.
	ErrorEntryNotExist = errors.New("entry not exist")

	// Any confict that should not have been occurred.
	ErrorConflict = errors.New("conflict, report immediately to the admin with record of behaviour")

	// use when bind json failed.
	ErrorBindJSON = errors.New("binding input failed")

	ErrorContextLostValue = errors.New("authentication failed during processing, please login again")

	ErrorNoParamReceived = errors.New("no params have been received by the server")
)

// Sql error
var (
	ErrorDeleteFailed = errors.New("deleting failed, please check the action again or report this to the admin with the record of behaviour")

	ErrorUpdateFailed = errors.New("updaing failed, please check the action again or report this to the admin with the record of behaviour")

	ErrorCreateFailed = errors.New("createing failed, please check the action again or report this to the admin with the record of behaviour")

	ErrorQueryFailed = errors.New("querying failed, please check the action again or report this to the admin with the record of behaviour")
)
