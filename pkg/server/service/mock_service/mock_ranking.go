// Code generated by MockGen. DO NOT EDIT.
// Source: ranking.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	service "InfecShotAPI/pkg/server/service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRankingServiceInterface is a mock of RankingServiceInterface interface.
type MockRankingServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRankingServiceInterfaceMockRecorder
}

// MockRankingServiceInterfaceMockRecorder is the mock recorder for MockRankingServiceInterface.
type MockRankingServiceInterfaceMockRecorder struct {
	mock *MockRankingServiceInterface
}

// NewMockRankingServiceInterface creates a new mock instance.
func NewMockRankingServiceInterface(ctrl *gomock.Controller) *MockRankingServiceInterface {
	mock := &MockRankingServiceInterface{ctrl: ctrl}
	mock.recorder = &MockRankingServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRankingServiceInterface) EXPECT() *MockRankingServiceInterfaceMockRecorder {
	return m.recorder
}

// GetRankInfoList mocks base method.
func (m *MockRankingServiceInterface) GetRankInfoList(serviceRequest *service.GetRankInfoListRequest) (*service.GetRankInfoListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRankInfoList", serviceRequest)
	ret0, _ := ret[0].(*service.GetRankInfoListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRankInfoList indicates an expected call of GetRankInfoList.
func (mr *MockRankingServiceInterfaceMockRecorder) GetRankInfoList(serviceRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRankInfoList", reflect.TypeOf((*MockRankingServiceInterface)(nil).GetRankInfoList), serviceRequest)
}
