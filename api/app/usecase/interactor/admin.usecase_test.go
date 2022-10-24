package interactor

import (
	"context"
	"server/app/domain/repository"
	"testing"

	"server/app/usecase/usecase_dto"
	mockusecase "server/app/usecase/usecase_mock"

	"github.com/golang/mock/gomock"
)

func buildAdminTest() *AdminUsecase {
	var repo repository.DataService
	adminusecase := NewAdminUseCase(repo)
	return adminusecase
}

func TestUpdateUserTestResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockAdminUseCase(mockCtrl)
	adminusecase := buildAdminTest()

	t.Run("Update Test Result", func(t *testing.T) {
		mockRepo.EXPECT().UpdateUserTestResult(context.Background(), usecase_dto.TestResult{}).Return(nil)
		err := adminusecase.UpdateUserTestResult(context.Background(), usecase_dto.TestResult{})
		if err != nil {
			t.Errorf("UpdateTestResult() error = %v", err)
		}
	})
}

func TestDeleteUserTestResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockAdminUseCase(mockCtrl)
	adminusecase := buildAdminTest()

	t.Run("Delete User Test Result", func(t *testing.T) {
		mockRepo.EXPECT().DeleteUserTestResult(context.Background(), 1).Return(nil)
		err := adminusecase.DeleteUserTestResult(context.Background(), 1)
		if err != nil {
			t.Errorf("DeleteUserTestResult() error = %v", err)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockAdminUseCase(mockCtrl)
	adminusecase := buildAdminTest()

	t.Run("Delete User", func(t *testing.T) {
		mockRepo.EXPECT().DeleteUser(context.Background(), 1).Return(nil)
		err := adminusecase.DeleteUser(context.Background(), 1)
		if err != nil {
			t.Errorf("DeleteUser() error = %v", err)
		}
	})
}
