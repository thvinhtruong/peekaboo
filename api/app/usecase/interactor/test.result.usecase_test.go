package interactor

import (
	"context"
	"reflect"
	"testing"

	"server/app/domain/repository"
	"server/app/usecase/usecase_dto"
	mockusecase "server/app/usecase/usecase_mock"

	"github.com/golang/mock/gomock"
)

func buildTestResultTest() *TestResultUsecase {
	var repo repository.DataService
	testresultusecase := NewTestResultUsecase(repo)
	return testresultusecase
}

func TestGetUserTestResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockTestResultUseCase(mockCtrl)
	testresultusecase := buildTestResultTest()

	t.Run("Get User Test Result", func(t *testing.T) {
		mockRepo.EXPECT().GetUserTestResults(context.Background(), 1).Return([]usecase_dto.TestResult{}, nil)
		_, err := testresultusecase.GetUserTestResults(context.Background(), 1)
		if err != nil {
			t.Errorf("GetUserTestResults() error = %v", err)
		}
	})
}

func TestGetUserTestResultDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockTestResultUseCase(mockCtrl)
	testresultusecase := buildTestResultTest()

	t.Run("Get User Test Result Detail", func(t *testing.T) {
		mockRepo.EXPECT().GetUserTestResultDetail(context.Background(), 1).Return(usecase_dto.TestResult{}, nil)
		got, err := testresultusecase.GetUserTestResultDetail(context.Background(), 1)
		if err != nil {
			t.Errorf("GetUserTestResultDetail() error = %v", err)
		}

		if !reflect.DeepEqual(got, usecase_dto.TestResult{}) {
			t.Errorf("GetUserTestResultDetail() got = %v", got)
		}
	})

}

func TestGetTestResultHeadline(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockTestResultUseCase(mockCtrl)
	testresultusecase := buildTestResultTest()

	t.Run("Get Test Result Headline", func(t *testing.T) {
		mockRepo.EXPECT().GetTestResultHeadline(context.Background(), 1).Return(usecase_dto.TestResult{}, nil)
		got, err := testresultusecase.GetTestResultHeadline(context.Background(), 1)
		if err != nil {
			t.Errorf("GetTestResultHeadline() error = %v", err)
		}

		if !reflect.DeepEqual(got, usecase_dto.TestResult{}) {
			t.Errorf("GetTestResultHeadline() got = %v", got)
		}
	})
}
