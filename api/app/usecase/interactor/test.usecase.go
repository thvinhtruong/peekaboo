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

type TestUsecase struct {
	TestRepository repository.DataService
}

func NewTestUsecase(testRepository repository.DataService) *TestUsecase {
	return &TestUsecase{
		TestRepository: testRepository,
	}
}

// Return only the test information.
func (t *TestUsecase) QueryTestInfo(ctx context.Context, testId int, userId int) (test usecase_dto.Test, err error) {
	testResultID := 0
	var dateCreated int64
	dateCreated = 0

	// check if user has done the task
	classes, err := t.TestRepository.QueryClassOfUser(ctx, userId)
	if err != nil {
		return test, err
	}

	for _, classId := range classes {
		tests, err := t.TestRepository.QueryTestOfClass(ctx, classId)
		if err != nil {
			return test, err
		}

		for _, tid := range tests {
			if tid.TestID == testId {
				testResult, err := t.TestRepository.QueryTestResultIndexScore(ctx, tid.ID, userId, time.Now(), 4)
				if err != nil {
					return test, err
				}

				if len(testResult) > 0 {
					for _, r := range testResult {
						if r.DateCreated > (dateCreated) {
							testResultID = r.ID
							dateCreated = r.DateCreated
						}
					}
				}
			}
		}
	}

	record, err := t.TestRepository.QueryTestHeadline(ctx, testId, "")
	if err != nil {
		return
	}

	if len(record) == 0 {
		return
	}

	if err := copier.Copy(&test, &record[0]); err != nil {
		return test, err
	}

	test.PreviousTestResultID = testResultID
	if test.PreviousTestResultID != 0 {
		test.IsDone = true
	}

	return
}

func (t *TestUsecase) QuerySkillTest(ctx context.Context, testId int) (test usecase_dto.SkillTest, err error) {
	skID, err := t.TestRepository.QuerySkillTestOfTest(ctx, testId)
	if err != nil {
		return
	}

	for _, v := range skID {
		record, err := t.TestRepository.QuerySkillTest(ctx, v)
		if err != nil {
			return test, err
		}

		if err := copier.Copy(&test, record); err != nil {
			return test, err
		}

		test.ID = v
	}

	return
}

// @transaction
// Steps:
// @1. Compare result to the database answer to produce the test result.
// @2. Insert the test answer into database.
func (t *TestUsecase) SubmitTest(ctx context.Context, data usecase_dto.SubmitData, userId int, entityCode int) (testResultId int, err error) {
	sk, err := t.TestRepository.QuerySkillTest(ctx, data.ID)
	if err != nil {
		return testResultId, err
	}

	if len(data.Sections) != len(sk.Section) {
		return testResultId, fmt.Errorf("the number of sections is not equal")

	}

	correctAns := 0
	totalAns := 0
	for i, section := range sk.Section {
		// @1. Compare result to the database answer to produce the test result.
		// @2. Save-up user's answer.
		if len(section.Content) != len(data.Sections[i].Answers) {
			return 0, fmt.Errorf("the number of sections is not equal")
		}

		for j, content := range section.Content {
			totalAns++
			if content.CorrectAns == data.Sections[i].Answers[j] {
				correctAns++
			}
		}
	}

	if err := t.TestRepository.EnableTx(func() error {
		// @1. Create test result
		testResultId, err = t.TestRepository.CreateTestResult(ctx, entity.TestResult{
			ID:          0,
			TestClassID: data.TestClassID,
			UserID:      userId,
			EntityCode:  entityCode,
			Score:       int(float32(correctAns) / float32(totalAns) * 100),
			Comment:     "N/A",
			ResultNote:  "Little Sloww",
		})

		if err != nil {
			return err
		}

		var entitySubmittedData entity.SubmittedAnswer
		if err := copier.Copy(&entitySubmittedData.Sections, &data.Sections); err != nil {
			return err
		}

		entitySubmittedData.ID = testResultId

		// @2. Insert the test answer into database.
		err = t.TestRepository.CreateTestAnswer(ctx, entitySubmittedData)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return testResultId, err
	}

	return testResultId, err
}

func (t *TestUsecase) QueryAllTest(ctx context.Context) (testResult []usecase_dto.Test, err error) {
	record, err := t.TestRepository.QueryAllTest(ctx)
	if err != nil {
		return
	}

	if err := copier.Copy(&testResult, &record); err != nil {
		return testResult, err
	}

	return
}

func (t *TestUsecase) QueryTestAnswer(ctx context.Context, resultId int) (testResult usecase_dto.SubmitData, err error) {
	record, err := t.TestRepository.FindTestAnswer(ctx, resultId)
	if err != nil {
		return
	}

	if err := copier.Copy(&testResult, &record); err != nil {
		return testResult, err
	}

	return testResult, nil
}
