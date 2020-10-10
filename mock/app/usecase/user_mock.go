// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/user.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	reflect "reflect"
)

// MockUserUsecaseInterface is a mock of UserUsecaseInterface interface
type MockUserUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseInterfaceMockRecorder
}

// MockUserUsecaseInterfaceMockRecorder is the mock recorder for MockUserUsecaseInterface
type MockUserUsecaseInterfaceMockRecorder struct {
	mock *MockUserUsecaseInterface
}

// NewMockUserUsecaseInterface creates a new mock instance
func NewMockUserUsecaseInterface(ctrl *gomock.Controller) *MockUserUsecaseInterface {
	mock := &MockUserUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUsecaseInterface) EXPECT() *MockUserUsecaseInterfaceMockRecorder {
	return m.recorder
}

// Authorization mocks base method
func (m *MockUserUsecaseInterface) Authorization(ctx context.Context, uname, password string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorization", ctx, uname, password)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorization indicates an expected call of Authorization
func (mr *MockUserUsecaseInterfaceMockRecorder) Authorization(ctx, uname, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorization", reflect.TypeOf((*MockUserUsecaseInterface)(nil).Authorization), ctx, uname, password)
}

// Register mocks base method
func (m *MockUserUsecaseInterface) Register(ctx context.Context, details model.SignUpDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, details)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUserUsecaseInterfaceMockRecorder) Register(ctx, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserUsecaseInterface)(nil).Register), ctx, details)
}