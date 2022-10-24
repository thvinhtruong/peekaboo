package mockusecase

import (
	"context"
	"reflect"
	"server/app/usecase/usecase_dto"

	"github.com/golang/mock/gomock"
)

type MockAdminUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAdminUseCaseMockRecorder
}

type MockAdminUseCaseMockRecorder struct {
	mock *MockAdminUseCase
}

func NewMockAdminUseCase(ctrl *gomock.Controller) *MockAdminUseCase {
	mock := &MockAdminUseCase{ctrl: ctrl}
	mock.recorder = &MockAdminUseCaseMockRecorder{mock}
	return mock
}

func (m *MockAdminUseCase) EXPECT() *MockAdminUseCaseMockRecorder {
	return m.recorder
}

func (m *MockAdminUseCase) UpdateUserTestResult(ctx context.Context, testResult usecase_dto.TestResult) (err error) {
	ret := m.ctrl.Call(m, "UpdateUserTestResult", ctx, testResult)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockAdminUseCaseMockRecorder) UpdateUserTestResult(ctx context.Context, testResult interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "UpdateUserTestResult", reflect.TypeOf((*MockAdminUseCase)(nil).UpdateUserTestResult), ctx, testResult)
}

func (m *MockAdminUseCase) DeleteUserTestResult(ctx context.Context, id int) (err error) {
	ret := m.ctrl.Call(m, "DeleteUserTestResult", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockAdminUseCaseMockRecorder) DeleteUserTestResult(ctx context.Context, id interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "DeleteUserTestResult", reflect.TypeOf((*MockAdminUseCase)(nil).DeleteUserTestResult), ctx, id)
}

func (m *MockAdminUseCase) DeleteUser(ctx context.Context, id int) (err error) {
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockAdminUseCaseMockRecorder) DeleteUser(ctx context.Context, id interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "DeleteUser", reflect.TypeOf((*MockAdminUseCase)(nil).DeleteUser), ctx, id)
}
