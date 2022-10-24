package interactor

import (
	"context"
	"server/app/domain/repository"
	"server/app/usecase/usecase_dto"
	"time"

	"github.com/jinzhu/copier"
)

type TestResultUsecase struct {
	TestResultRepository repository.DataService
}

func NewTestResultUsecase(TestSkillRepository repository.DataService) *TestResultUsecase {
	return &TestResultUsecase{
		TestResultRepository: TestSkillRepository,
	}
}

// Only return the score of each test result.
// @deprecated
func (s *TestResultUsecase) GetUserTestResults(ctx context.Context, userID int) (results []usecase_dto.TestResult, err error) {
	classes, err := s.TestResultRepository.QueryClassOfUser(ctx, userID)
	if err != nil {
		return
	}

	for _, v := range classes {
		testClass, err := s.TestResultRepository.QueryTestOfClass(ctx, v)
		if err != nil {
			return nil, err
		}

		for _, tc := range testClass {
			testResult, err := s.TestResultRepository.QueryTestResultIndexScore(ctx, tc.ID, userID, time.Now(), 4)
			if err != nil {
				return nil, err
			}

			if len(testResult) == 0 {
				continue
			}

			var tr []usecase_dto.TestResult
			if err := copier.Copy(&tr, &testResult); err != nil {
				return nil, err
			}

			results = append(results, tr...)
		}
	}

	return
}

// Only the test result with information.
// @deprecated
func (s *TestResultUsecase) GetUserTestResultDetail(ctx context.Context, testId int) (result usecase_dto.TestResult, err error) {
	record, err := s.TestResultRepository.QueryTestResultDetails(ctx, testId)
	if err != nil {
		return result, err
	}

	if err := copier.Copy(&result, &record); err != nil {
		return result, err
	}

	return
}

func (s *TestResultUsecase) GetTestResultHeadline(ctx context.Context, testResultId int) (result usecase_dto.TestResult, err error) {
	record, err := s.TestResultRepository.QueryTestResultDetails(ctx, testResultId)
	if err != nil {
		return result, err
	}

	if len(record) == 0 {
		return result, nil
	}

	if err := copier.Copy(&result, &record[0]); err != nil {
		return result, err
	}

	return result, err
}
