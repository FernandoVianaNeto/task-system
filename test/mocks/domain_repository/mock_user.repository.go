// Code generated by MockGen. DO NOT EDIT.
// Source: user.repository.go
//
// Generated by this command:
//
//	mockgen -source user.repository.go -package domain_repository -destination /home/fernandoviana/Documentos/Pessoal/projetos/task-system/test/mocks/domain_repository/mock_user.repository.go
//

// Package domain_repository is a generated GoMock package.
package domain_repository

import (
	context "context"
	reflect "reflect"
	dto "task-system/internal/domain/dto"
	entities "task-system/internal/domain/entities"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepositoryInterface) CreateUser(ctx context.Context, input entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateUser(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateUser), ctx, input)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepositoryInterface) GetUserByEmail(ctx context.Context, input dto.GetUserByEmailDto) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, input)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUserByEmail(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUserByEmail), ctx, input)
}

// GetUserByUuid mocks base method.
func (m *MockUserRepositoryInterface) GetUserByUuid(ctx context.Context, input dto.GetUserByUuidDto) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUuid", ctx, input)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUuid indicates an expected call of GetUserByUuid.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUserByUuid(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUuid", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUserByUuid), ctx, input)
}
