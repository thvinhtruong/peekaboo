package usecase

import (
	"context"
	"server/app/usecase/usecase_dto"
)

type AdminService interface {
	UpdateUserTestResult(ctx context.Context, testResult usecase_dto.TestResult) (err error)

	DeleteUserTestResult(ctx context.Context, id int) (err error)

	// @transaction
	DeleteUser(ctx context.Context, id int) (err error)
}

type UserService interface {
	FindAllUser(ctx context.Context) (result []usecase_dto.User, err error)

	CreateUser(ctx context.Context, user usecase_dto.User) (id int, err error)

	UpdateUser(ctx context.Context, user usecase_dto.User) (err error)

	FindUser(ctx context.Context, user usecase_dto.User, HasPassword bool) (result []usecase_dto.User, err error)

	ReviewTestResult(ctx context.Context, resultId int) (submittedAnswer usecase_dto.SubmitData, err error)

	FindUserClasses(ctx context.Context, userId int) (classes []usecase_dto.Class, err error)

	FindAllUserTestResult(ctx context.Context, userId int) (results []usecase_dto.TestResult, err error)
}

type ClassService interface {
	DeleteClass(ctx context.Context, classId int) error

	CreateClass(ctx context.Context, class usecase_dto.Class) (int, error)

	GetClasses(ctx context.Context) (classes []usecase_dto.Class, err error)

	QueryClassMembers(ctx context.Context, classId int) (users []usecase_dto.User, err error)

	AddMember2Class(ctx context.Context, classId int, userId int) (err error)

	RemoveMemberFromClass(ctx context.Context, classId int, userId int) (err error)

	QueryClassTestResult(ctx context.Context, classId int, testId int) (results []usecase_dto.TestResult, err error)

	GetClassTest(ctx context.Context, classId int) (tests []usecase_dto.Test, err error)

	AddTest2Class(ctx context.Context, classId int, testId int) (err error)

	RemoveTestClass(ctx context.Context, classId int, testId int) (err error)

	GetSingleClassTest(ctx context.Context, classId int, testId int) (test usecase_dto.Test, err error)
}

type TestService interface {
	// GetTestByID returns a test by its ID.
	QueryTestInfo(ctx context.Context, testId int, userId int) (test usecase_dto.Test, err error)

	QuerySkillTest(ctx context.Context, testId int) (test usecase_dto.SkillTest, err error)

	SubmitTest(ctx context.Context, data usecase_dto.SubmitData, userId int, entityCode int) (testResultId int, err error)

	QueryAllTest(ctx context.Context) (testResult []usecase_dto.Test, err error)

	QueryTestAnswer(ctx context.Context, resultId int) (testResult usecase_dto.SubmitData, err error)
}

type TestResultService interface {
	GetUserTestResults(ctx context.Context, userID int) (results []usecase_dto.TestResult, err error)

	GetUserTestResultDetail(ctx context.Context, testId int) (result usecase_dto.TestResult, err error)

	GetTestResultHeadline(ctx context.Context, testResultId int) (result usecase_dto.TestResult, err error)
}

type TestSkillService interface {
	QueryTestResult(ctx context.Context, userId int)

	GetAllUserTestResults(ctx context.Context, userID int) (results []usecase_dto.TestResult, err error)
}
