// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	model "InfecShotAPI/pkg/server/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// InsertUser mocks base method.
func (m *MockUserRepositoryInterface) InsertUser(record *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) InsertUser(record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).InsertUser), record)
}

// SelectUserByAuthToken mocks base method.
func (m *MockUserRepositoryInterface) SelectUserByAuthToken(authToken string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserByAuthToken", authToken)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserByAuthToken indicates an expected call of SelectUserByAuthToken.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectUserByAuthToken(authToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserByAuthToken", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectUserByAuthToken), authToken)
}

// SelectUserByPrimaryKey mocks base method.
func (m *MockUserRepositoryInterface) SelectUserByPrimaryKey(userID string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserByPrimaryKey", userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserByPrimaryKey indicates an expected call of SelectUserByPrimaryKey.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectUserByPrimaryKey(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserByPrimaryKey", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectUserByPrimaryKey), userID)
}

// SelectUsersOrderByHighScoreAsc mocks base method.
func (m *MockUserRepositoryInterface) SelectUsersOrderByHighScoreAsc(limit, offset int) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUsersOrderByHighScoreAsc", limit, offset)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUsersOrderByHighScoreAsc indicates an expected call of SelectUsersOrderByHighScoreAsc.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectUsersOrderByHighScoreAsc(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUsersOrderByHighScoreAsc", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectUsersOrderByHighScoreAsc), limit, offset)
}

// UpdateUserByPrimaryKey mocks base method.
func (m *MockUserRepositoryInterface) UpdateUserByPrimaryKey(record *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByPrimaryKey", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByPrimaryKey indicates an expected call of UpdateUserByPrimaryKey.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdateUserByPrimaryKey(record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByPrimaryKey", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdateUserByPrimaryKey), record)
}