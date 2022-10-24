package interactor

import (
	"context"
	"server/app/domain/entity"
	"server/app/domain/repository"
	"server/app/usecase/usecase_dto"
	"time"

	"github.com/jinzhu/copier"
)

type AdminUsecase struct {
	adminRepo repository.DataService
}

func NewAdminUseCase(adminRepo repository.DataService) *AdminUsecase {
	return &AdminUsecase{
		adminRepo: adminRepo,
	}
}

func (a *AdminUsecase) UpdateUserTestResult(ctx context.Context, testResult usecase_dto.TestResult) (err error) {
	var testResultEntity entity.TestResult
	if err := copier.Copy(&testResultEntity, &testResult); err != nil {
		return err
	}

	return a.adminRepo.UpdateTestResult(ctx, testResultEntity)
}

func (a *AdminUsecase) DeleteUserTestResult(ctx context.Context, id int) (err error) {
	if err := a.adminRepo.DeleteTestResult(ctx, id); err != nil {
		return err
	}

	return nil
}

// Delete user will:
// @transaction
// 1. Remove user from all classes.
// 2. Remove all user's test result inside their class.
// 3. Remove user.
func (a *AdminUsecase) DeleteUser(ctx context.Context, id int) (err error) {
	classes, err := a.adminRepo.QueryClassOfUser(ctx, id)
	if err != nil {
		return err
	}

	return a.adminRepo.EnableTx(func() error {
		// go through classes
		for _, v := range classes {
			// test result
			testClass, err := a.adminRepo.QueryTestOfClass(ctx, v)
			if err != nil {
				return err
			}

			for _, tc := range testClass {
				testResult, err := a.adminRepo.QueryTestResultIndexScore(ctx, tc.ID, id, time.Now(), 4)
				if err != nil {
					return err
				}

				if len(testResult) == 0 {
					continue
				}

				for _, r := range testResult {
					// delete the user's test results inside the classes.
					if err := a.adminRepo.DeleteTestResult(ctx, r.ID); err != nil {
						return err
					}
				}
			}

			if err := a.adminRepo.DeleteUserClass(ctx, v, id); err != nil {
				return err
			}
		}

		// delete user
		if err := a.adminRepo.DeleteUser(ctx, id); err != nil {
			return err
		}

		return nil
	})
}
