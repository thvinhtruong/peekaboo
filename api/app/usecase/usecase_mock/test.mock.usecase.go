package mockusecase

import (
	"context"
	"reflect"
	"server/app/usecase/usecase_dto"

	"github.com/golang/mock/gomock"
)

type MockTestUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTestUseCaseMockRecorder
}

type MockTestUseCaseMockRecorder struct {
	mock *MockTestUseCase
}

func NewMockTestUseCase(ctrl *gomock.Controller) *MockTestUseCase {
	mock := &MockTestUseCase{ctrl: ctrl}
	mock.recorder = &MockTestUseCaseMockRecorder{mock}
	return mock
}

func (m *MockTestUseCase) EXPECT() *MockTestUseCaseMockRecorder {
	return m.recorder
}

func (m *MockTestUseCase) QueryTestInfo(ctx context.Context, testId int) (test usecase_dto.Test, err error) {
	ret := m.ctrl.Call(m, "QueryTestInfo", ctx, testId)
	ret0, _ := ret[0].(usecase_dto.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTestUseCaseMockRecorder) QueryTestInfo(ctx context.Context, testId interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "QueryTestInfo", reflect.TypeOf((*MockTestUseCase)(nil).QueryTestInfo))
}

func (m *MockTestUseCase) QuerySkillTest(ctx context.Context, testId int) (test usecase_dto.SkillTest, err error) {
	ret := m.ctrl.Call(m, "QuerySkillTest", ctx, testId)
	ret0, _ := ret[0].(usecase_dto.SkillTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTestUseCaseMockRecorder) QuerySkillTest(ctx context.Context, testId interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "QuerySkillTest", reflect.TypeOf((*MockTestUseCase)(nil).QuerySkillTest))
}

func (m *MockTestUseCase) SubmitTest(ctx context.Context, data usecase_dto.SubmitData, userId int, entityCode int) (testResultId int, err error) {
	ret := m.ctrl.Call(m, "SubmitTest", ctx, data, userId, entityCode)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTestUseCaseMockRecorder) SubmitTest(ctx context.Context, data usecase_dto.SubmitData, userId interface{}, entityCode interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "SubmitTest", reflect.TypeOf((*MockTestUseCase)(nil).SubmitTest))
}
