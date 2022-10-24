package repository

import (
	"context"
	"server/app/domain/entity"
	"time"
)

type DataService interface {
	UserRepository
	TestRepository
	TestResultRepository
	SkillTestRepository
	ClassRepository
	TagRepository

	TransactionMng
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)

	// Flag determine the query element of the functions: [0 - username]  [1 - fullname] [2 - ID].
	// Option 2 is unindex search, therefore do not overuse.
	QueryUser(ctx context.Context, Username string, Fullname string, ID int, Flag int, HasPassword bool) ([]entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, ID int) error
}

type TestRepository interface {
	QueryAllTest(ctx context.Context) (result []entity.Test, err error)
	CreateTest(ctx context.Context, test entity.Test) (int, error)
	QueryTestHeadline(ctx context.Context, ID int, TestName string) (result []entity.Test, err error)
	QueryTestByTestTag(ctx context.Context, TestTagID int) (result []entity.Test, err error)
	QueryAllTestHeadlines(ctx context.Context) ([]entity.Test, error)
	UpdateTest(ctx context.Context, test entity.Test) error
	DeleteTest(ctx context.Context, TestID int) error

	AssignSkillTest2Test(ctx context.Context, TestID int, SkillTestID int) error
	QuerySkillTestOfTest(ctx context.Context, TestID int) (result []int, err error)
	DeleteSkillTestAndTest(ctx context.Context, TestID int, SkillTestID int) error

	AssignTestClass(ctx context.Context, TestClass entity.TestClassRelation) error
	QueryTestClass(ctx context.Context, TestClassID int) (result []entity.TestClassRelation, err error)
	QueryTestOfClass(ctx context.Context, ClassID int) (result []entity.TestClassRelation, err error)
	QueryClassDoneTest(ctx context.Context, TestID int) (result []entity.TestClassRelation, err error)
	DeleteTestClass(ctx context.Context, TestID int, ClassID int) error

	UnarchieveUserClass(ctx context.Context, UserID int, ClassID int) (err error)
	CheckExistedUserClass(ctx context.Context, UserID int, ClassID int) (existed bool, err error)
	AddUserClass(ctx context.Context, ClassID int, UserID int) error
	QueryUserOfClass(ctx context.Context, ClassID int) ([]int, error)
	QueryClassOfUser(ctx context.Context, UserID int) ([]int, error)
	DeleteUserClass(ctx context.Context, ClassID int, StudentID int) error

	CreateTestAnswer(ctx context.Context, ans entity.SubmittedAnswer) (err error)
	DeleteTestAnswer(ctx context.Context, ans entity.SubmittedAnswer) (err error)
	UpdateTestAnswer(ctx context.Context, ans entity.SubmittedAnswer) (err error)
	FindTestAnswer(ctx context.Context, id int) (ans entity.SubmittedAnswer, err error)
}

type TestResultRepository interface {
	CreateTestResult(ctx context.Context, data entity.TestResult) (int, error)
	UpdateTestResult(ctx context.Context, data entity.TestResult) error
	QueryTestResultDetails(ctx context.Context, ID int) ([]entity.TestResult, error)

	// Query TestResultIndex allows you to search according to flag, but only return the indexed value: ID, TestID, user_id, ClassID. Other fields are expected to be 0/nil.
	// Flag determine the query element of the functions: [1 - TestClassID] [2 - user_id] [3 - DateCreated] [4 - TestClassID, user_id] [5 - TestClassID, DateCreated] [6 - user_id, DateCreated] [7 - TestClassID, user_id, DateCreated]
	QueryTestResultIndexScore(ctx context.Context, TestClassID int, UserID int, DateCreate time.Time, Flag int) ([]entity.TestResult, error)
	DeleteTestResult(ctx context.Context, ID int) error
	ArchieveTestResult(ctx context.Context, TestClassID int, UserID int) error
	DeleteTestResultOfClass(ctx context.Context, TestClassID int) error
}

type SkillTestRepository interface {
	CreateSkillTest(ctx context.Context, st entity.SkillTest) (id int, err error)
	QuerySkillTest(ctx context.Context, id int) (st entity.SkillTest, err error)
	UpdateSkillTest(ctx context.Context, st entity.SkillTest) (err error)
	DeleteSkillTest(ctx context.Context, st entity.SkillTest) (err error)
}

type ClassRepository interface {
	DeleteClass(ctx context.Context, ID int) error
	CreateClass(ctx context.Context, class entity.Class) (int, error)
	QueryAllClass(ctx context.Context) ([]entity.Class, error)
	QueryClass(ctx context.Context, ClassID int, ClassName string) ([]entity.Class, error)

	AddUserClass(ctx context.Context, ClassID int, UserID int) error
	QueryUserOfClass(ctx context.Context, ClassID int) ([]int, error)
	QueryClassOfUser(ctx context.Context, UserID int) ([]int, error)
	QueryTestOfClass(ctx context.Context, ClassID int) (result []entity.TestClassRelation, err error)
	DeleteUserClass(ctx context.Context, ClassID int, StudentID int) error
}

type TagRepository interface {
	CreateTag(ctx context.Context, testtag entity.Tag) (int, error)
	UpdateTag(ctx context.Context, testtag entity.Tag) error
	QueryTag(ctx context.Context, TestTagID int, flag int) ([]entity.Tag, error)
	DeleteTag(ctx context.Context, TestTagID int) error
}

type TransactionMng interface {
	EnableTx(txFunc func() error) error
}
