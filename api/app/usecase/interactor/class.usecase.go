package interactor

import (
	"context"
	"fmt"
	"server/app/domain/entity"
	"server/app/domain/repository"
	"server/app/usecase/usecase_dto"
	"time"

	"github.com/jinzhu/copier"
)

type ClassUsecase struct {
	ClassRepository repository.DataService
}

func NewClassUsecase(DataService repository.DataService) *ClassUsecase {
	return &ClassUsecase{
		ClassRepository: DataService,
	}
}

// @transaction
// Delete all the test result from test from the class.
// Delete all the test class id.
// Delete all the user class id.
// Delete the class id.
func (c *ClassUsecase) DeleteClass(ctx context.Context, classId int) error {
	TestClassID, err := c.ClassRepository.QueryTestOfClass(ctx, classId)
	if err != nil {
		return err
	}

	UserClassID, err := c.ClassRepository.QueryUserOfClass(ctx, classId)
	if err != nil {
		return err
	}

	return c.ClassRepository.EnableTx(func() error {
		for _, v := range TestClassID {
			testResult, err := c.ClassRepository.QueryTestResultIndexScore(ctx, v.ID, 0, time.Now(), 1)
			if err != nil {
				return err
			}

			for _, r := range testResult {
				err := c.ClassRepository.DeleteTestResult(ctx, r.ID)
				if err != nil {
					return err
				}
			}

			err = c.ClassRepository.DeleteTestClass(ctx, v.TestID, v.ClassID)
			if err != nil {
				return err
			}
		}

		for _, id := range UserClassID {
			if err := c.ClassRepository.DeleteUserClass(ctx, classId, id); err != nil {
				return err
			}
		}

		if err := c.ClassRepository.DeleteClass(ctx, classId); err != nil {
			return err
		}

		return nil
	})
}

func (c *ClassUsecase) CreateClass(ctx context.Context, class usecase_dto.Class) (int, error) {
	var record entity.Class
	err := copier.Copy(&record, &class)
	if err != nil {
		return 0, err
	}

	id, err := c.ClassRepository.CreateClass(ctx, record)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (c *ClassUsecase) GetClasses(ctx context.Context) (classes []usecase_dto.Class, err error) {
	records, err := c.ClassRepository.QueryAllClass(ctx)
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		var class usecase_dto.Class
		err = copier.Copy(&class, &record)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, nil
}

func (c *ClassUsecase) QueryClassMembers(ctx context.Context, classId int) (users []usecase_dto.User, err error) {
	records_id, err := c.ClassRepository.QueryUserOfClass(ctx, classId)
	if err != nil {
		return nil, err
	}

	for _, record_id := range records_id {
		var user usecase_dto.User
		userDB, err := c.ClassRepository.QueryUser(ctx, "", "", record_id, 2, false)
		if err != nil {
			return nil, err
		}

		if len(userDB) == 0 {
			continue
		}

		err = copier.Copy(&user, &userDB[0])
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (c *ClassUsecase) AddMember2Class(ctx context.Context, classId int, userId int) (err error) {
	existed, err := c.ClassRepository.CheckExistedUserClass(ctx, userId, classId)
	if err != nil {
		return err
	}

	switch existed {
	case true:
		err = c.ClassRepository.UnarchieveUserClass(ctx, userId, classId)
		if err != nil {
			return err
		}

		return
	default:
		err = c.ClassRepository.AddUserClass(ctx, classId, userId)
		if err != nil {
			return err
		}

		return
	}
}

func (c *ClassUsecase) AddTest2Class(ctx context.Context, classId int, testId int) (err error) {
	err = c.ClassRepository.AssignTestClass(ctx, entity.TestClassRelation{
		TestID:  testId,
		ClassID: classId,
	})

	if err != nil {
		return err
	}

	return

}

func (c *ClassUsecase) RemoveMemberFromClass(ctx context.Context, classId int, userId int) (err error) {
	err = c.ClassRepository.DeleteUserClass(ctx, classId, userId)
	if err != nil {
		return err
	}
	return nil
}

// @reduce
func (c *ClassUsecase) RemoveTestClass(ctx context.Context, classId int, testId int) (err error) {
	err = c.ClassRepository.DeleteTestClass(ctx, testId, classId)
	if err != nil {
		return err
	}

	return

}

// get all reults of corresponding class and test
func (c *ClassUsecase) QueryClassTestResult(ctx context.Context, classId int, testId int) (results []usecase_dto.TestResult, err error) {
	var testClassId int
	var totalRecord []entity.TestResult
	classDoneTest, err := c.ClassRepository.QueryClassDoneTest(ctx, testId)
	if err != nil {
		return results, err
	}

	for _, item := range classDoneTest {
		if item.ClassID == classId {
			testClassId = item.ID
		} else {
			return nil, fmt.Errorf("this class has not taken this test yet")
		}
	}

	userId, err := c.ClassRepository.QueryUserOfClass(ctx, classId) // list of UserId (int)
	if err != nil {
		return results, err
	}

	for _, item := range userId {
		if item != 0 {
			// date created will be ignored in this case
			// query followed by userId and testClassId
			records, err := c.ClassRepository.QueryTestResultIndexScore(ctx, testClassId, item, time.Unix(0, 0), 4)
			if err != nil {
				return results, err
			}
			totalRecord = append(totalRecord, records...)
		}
	}
	err = copier.Copy(&results, &totalRecord)
	if err != nil {
		return results, err
	}
	return results, nil

}

// get all test within a class
func (c *ClassUsecase) GetClassTest(ctx context.Context, classId int) (tests []usecase_dto.Test, err error) {
	var test usecase_dto.Test
	records, err := c.ClassRepository.QueryTestOfClass(ctx, classId)
	if err != nil {
		return nil, err
	}

	for _, item := range records {
		record_item, err := c.ClassRepository.QueryTestHeadline(ctx, item.TestID, "")
		if err != nil {
			return nil, err
		}

		if len(record_item) == 0 {
			continue
		}

		if err := copier.Copy(&test, &record_item[0]); err != nil {
			return nil, err
		}

		test.TestClassID = item.ID
		tests = append(tests, test)
	}

	return tests, nil
}

// get all test within a class
func (c *ClassUsecase) GetSingleClassTest(ctx context.Context, classId int, testId int) (test usecase_dto.Test, err error) {
	records, err := c.ClassRepository.QueryTestOfClass(ctx, classId)
	if err != nil {
		return test, err
	}

	for _, item := range records {
		if item.TestID == testId {
			record_item, err := c.ClassRepository.QueryTestHeadline(ctx, item.TestID, "")
			if err != nil {
				return test, err
			}

			if len(record_item) == 0 {
				continue
			}

			if err := copier.Copy(&test, &record_item[0]); err != nil {
				return test, err
			}

			test.TestClassID = item.ID

			return test, nil
		}

	}

	return test, nil
}
