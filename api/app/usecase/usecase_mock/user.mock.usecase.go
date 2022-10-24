package mockusecase

import (
	"context"
	"reflect"
	"server/app/usecase/usecase_dto"

	"github.com/golang/mock/gomock"
)

type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

func (m *MockUserUseCase) CreateUser(ctx context.Context, user usecase_dto.User) error {
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockUserUseCaseMockRecorder) CreateUser(ctx context.Context, user interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "CreateUser", reflect.TypeOf((*MockUserUseCase)(nil).CreateUser), ctx, user)
}

func (m *MockUserUseCase) UpdateUser(ctx context.Context, user usecase_dto.User) error {
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockUserUseCaseMockRecorder) UpdateUser(ctx context.Context, user interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "UpdateUser", reflect.TypeOf((*MockUserUseCase)(nil).UpdateUser), ctx, user)
}

func (m *MockUserUseCase) FindUser(ctx context.Context, user usecase_dto.User, HasPassword bool) (result []usecase_dto.User, err error) {
	ret := m.ctrl.Call(m, "FindUser", ctx, user, HasPassword)
	ret0, _ := ret[0].([]usecase_dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockUserUseCaseMockRecorder) FindUser(ctx context.Context, user interface{}, HasPassword interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "FindUser", reflect.TypeOf((*MockUserUseCase)(nil).FindUser), ctx, user, HasPassword)
}

func (m *MockUserUseCase) FindUserClasses(ctx context.Context, userId int) (classes []usecase_dto.Class, err error) {
	ret := m.ctrl.Call(m, "FindUserClasses", ctx, userId)
	ret0, _ := ret[0].([]usecase_dto.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockUserUseCaseMockRecorder) FindUserClasses(ctx context.Context, userId interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "FindUserClasses", reflect.TypeOf((*MockUserUseCase)(nil).FindUserClasses), ctx, userId)
}

func (m *MockUserUseCase) ReviewTestResult(ctx context.Context, resultId int) (skilltest usecase_dto.SkillTest, err error) {
	ret := m.ctrl.Call(m, "ReviewTestResult", ctx, resultId)
	ret0, _ := ret[0].(usecase_dto.SkillTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockUserUseCaseMockRecorder) ReviewTestResult(ctx context.Context, resultId interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "ReviewTestResult", reflect.TypeOf((*MockUserUseCase)(nil).ReviewTestResult), ctx, resultId)
}
