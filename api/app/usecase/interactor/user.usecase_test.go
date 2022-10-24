package interactor

import (
	"context"
	"errors"
	"reflect"
	"server/app/domain/repository"
	"server/app/usecase/usecase_dto"
	mockusecase "server/app/usecase/usecase_mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func buildUserTest() *UserUsecase {
	var repo repository.DataService
	userusecase := NewUserUsecase(repo)
	return userusecase
}

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user usecase_dto.User
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "CreateUserSuccess",
			args: args{
				ctx:  context.Background(),
				user: usecase_dto.User{},
			},
			want:    int(0),
			wantErr: nil,
		},

		{
			name: "CreateUserError",
			args: args{
				ctx:  context.Background(),
				user: usecase_dto.User{},
			},
			want:    int(0),
			wantErr: nil,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockUserUseCase(mockCtrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "CreateUserSuccess" {
				mockRepo.EXPECT().CreateUser(tt.args.ctx, tt.args.user).Return(int(1), nil)
				userusecase := buildUserTest()
				got, err := userusecase.CreateUser(tt.args.ctx, tt.args.user)
				if err != nil {
					t.Errorf("UserUsecase.CreateUser() error = %v", err)
				}

				if got != tt.want {
					t.Errorf("CreateUser() = %v, want %v", got, tt.want)
				}
			}
			if tt.name == "CreateUserError" {
				mockRepo.EXPECT().CreateUser(tt.args.ctx, tt.args.user).Return(int(0), tt.wantErr)
				userusecase := buildUserTest()
				got, err := userusecase.CreateUser(tt.args.ctx, tt.args.user)
				if err != tt.wantErr {
					t.Errorf("UserUsecase.CreateUser() error = %v, want %v", err, tt.wantErr)
				}

				if got != tt.want {
					t.Errorf("CreateUser() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user usecase_dto.User
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "UpdateUserSuccess",
			args: args{
				ctx:  context.Background(),
				user: usecase_dto.User{},
			},
			want: nil,
		},
		{
			name: "UpdateUserError",
			args: args{
				ctx:  context.Background(),
				user: usecase_dto.User{},
			},
			want: errors.New("can not update user"),
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockUserUseCase(mockCtrl)

	for _, tt := range tests {
		if tt.name == "UpdateUserSuccess" {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.EXPECT().UpdateUser(tt.args.ctx, tt.args.user).Return(nil)
				userusecase := buildUserTest()
				if got := userusecase.UpdateUser(tt.args.ctx, tt.args.user); got != tt.want {
					t.Errorf("UserUsecase.UpdateUser() = %v error = %v", tt.want, got)
				}
			})
		}
		if tt.name == "UpdateUserError" {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.EXPECT().UpdateUser(tt.args.ctx, tt.args.user).Return(errors.New("can not update user"))
				userusecase := buildUserTest()
				got := userusecase.UpdateUser(tt.args.ctx, tt.args.user)
				if got != tt.want {
					t.Errorf("UserUsecase.UpdateUser() error = %v", got)
				}
			})
		}
	}
}

func TestFindUser(t *testing.T) {
	type args struct {
		ctx         context.Context
		user        usecase_dto.User
		hasPassword bool
	}

	tests := []struct {
		name    string
		args    args
		want    []usecase_dto.User
		wantErr error
	}{
		{
			name: "FindUserSuccess",
			args: args{
				ctx:         context.Background(),
				user:        usecase_dto.User{},
				hasPassword: true,
			},
			want:    []usecase_dto.User{},
			wantErr: nil,
		},

		{
			name: "FindUserError",
			args: args{
				ctx:         context.Background(),
				user:        usecase_dto.User{},
				hasPassword: true,
			},
			want:    []usecase_dto.User{},
			wantErr: errors.New("can not find user"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockUserUseCase(mockCtrl)

	for _, tt := range tests {
		if tt.name == "FindUserSuccess" {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.EXPECT().FindUser(tt.args.ctx, tt.args.user, tt.args.hasPassword).Return(tt.want, nil)
				userusecase := buildUserTest()
				got, err := userusecase.FindUser(tt.args.ctx, tt.args.user, tt.args.hasPassword)
				if err != nil {
					t.Errorf("UserUsecase.FindUser() error = %v", err)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FindUser() = %v, want %v", got, tt.want)
				}
			})
		}

		if tt.name == "FindUserError" {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.EXPECT().FindUser(tt.args.ctx, tt.args.user, tt.args.hasPassword).Return(tt.want, tt.wantErr)
				userusecase := buildUserTest()
				got, err := userusecase.FindUser(tt.args.ctx, tt.args.user, tt.args.hasPassword)
				if err != tt.wantErr {
					t.Errorf("UserUsecase.FindUser() error = %v, want %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FindUser() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestFindUserClasses(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId int
	}

	tests := []struct {
		name    string
		args    args
		want    []usecase_dto.Class
		wantErr error
	}{
		{
			name: "FindUserClassesSuccess",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
			want:    []usecase_dto.Class{},
			wantErr: nil,
		},

		{
			name: "FindUserClassesError",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
			want:    []usecase_dto.Class{},
			wantErr: errors.New("can not find user classes"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockUserUseCase(mockCtrl)

	for _, tt := range tests {
		if tt.name == "FindUserClassesSuccess" {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.EXPECT().FindUserClasses(tt.args.ctx, tt.args.userId).Return(tt.want, nil)
				userusecase := buildUserTest()
				got, err := userusecase.FindUserClasses(tt.args.ctx, tt.args.userId)
				if err != nil {
					t.Errorf("UserUsecase.FindUserClasses() error = %v", err)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FindUserClasses() = %v, want %v", got, tt.want)
				}
			})
		}

		if tt.name == "FindUserClassesError" {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.EXPECT().FindUserClasses(tt.args.ctx, tt.args.userId).Return(tt.want, tt.wantErr)
				userusecase := buildUserTest()
				got, err := userusecase.FindUserClasses(tt.args.ctx, tt.args.userId)
				if err != tt.wantErr {
					t.Errorf("UserUsecase.FindUserClasses() error = %v, want %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FindUserClasses() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestReviewTestResults(t *testing.T) {
	type args struct {
		ctx      context.Context
		resultId int
	}

	tests := []struct {
		name    string
		args    args
		want    usecase_dto.SkillTest
		wantErr error
	}{
		{
			name: "ReviewTestResults",
			args: args{
				ctx:      context.Background(),
				resultId: 1,
			},
			want:    usecase_dto.SkillTest{},
			wantErr: nil,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockusecase.NewMockUserUseCase(mockCtrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().ReviewTestResult(tt.args.ctx, tt.args.resultId).Return(tt.want, tt.wantErr)
			userusecase := buildUserTest()
			got, err := userusecase.ReviewTestResult(tt.args.ctx, tt.args.resultId)

			if err != tt.wantErr {
				t.Errorf("UserUsecase.ReviewTestResults() error = %v, want %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReviewTestResults() = %v, want %v", got, tt.want)
			}
		})
	}

}


