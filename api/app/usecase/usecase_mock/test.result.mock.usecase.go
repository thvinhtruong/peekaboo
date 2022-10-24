package mockusecase

import (
	"context"
	"reflect"
	"server/app/usecase/usecase_dto"

	"github.com/golang/mock/gomock"
)

type MockTestResultUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTestResultUseCaseMockRecorder
}

type MockTestResultUseCaseMockRecorder struct {
	mock *MockTestResultUseCase
}

func NewMockTestResultUseCase(ctrl *gomock.Controller) *MockTestResultUseCase {
	mock := &MockTestResultUseCase{ctrl: ctrl}
	mock.recorder = &MockTestResultUseCaseMockRecorder{mock}
	return mock
}

func (m *MockTestResultUseCase) EXPECT() *MockTestResultUseCaseMockRecorder {
	return m.recorder
}

func (m *MockTestResultUseCase) GetUserTestResults(ctx context.Context, userID int) (results []usecase_dto.TestResult, err error) {
	ret := m.ctrl.Call(m, "GetUserTestResults", ctx, userID)
	ret0, _ := ret[0].([]usecase_dto.TestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTestResultUseCaseMockRecorder) GetUserTestResults(ctx context.Context, userID interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "GetUserTestResults", reflect.TypeOf((*MockTestResultUseCase)(nil).GetUserTestResults), ctx, userID)
}

func (m *MockTestResultUseCase) GetUserTestResultDetail(ctx context.Context, testId int) (result usecase_dto.TestResult, err error) {
	ret := m.ctrl.Call(m, "GetUserTestResultDetail", ctx, testId)
	ret0, _ := ret[0].(usecase_dto.TestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTestResultUseCaseMockRecorder) GetUserTestResultDetail(ctx context.Context, testId interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "GetUserTestResultDetail", reflect.TypeOf((*MockTestResultUseCase)(nil).GetUserTestResultDetail), ctx, testId)
}

func (m *MockTestResultUseCase) GetTestResultHeadline(ctx context.Context, testResultId int) (result usecase_dto.TestResult, err error) {
	ret := m.ctrl.Call(m, "GetTestResultHeadline", ctx, testResultId)
	ret0, _ := ret[0].(usecase_dto.TestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTestResultUseCaseMockRecorder) GetTestResultHeadline(ctx context.Context, testResultId interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "GetTestResultHeadline", reflect.TypeOf((*MockTestResultUseCase)(nil).GetTestResultHeadline), ctx, testResultId)
}
