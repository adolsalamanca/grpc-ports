// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/port.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	"github.com/adolsalamanca/grpc-ports/server/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockPortRepository is a mock of PortRepository interface
type MockPortRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPortRepositoryMockRecorder
}

// MockPortRepositoryMockRecorder is the mock recorder for MockPortRepository
type MockPortRepositoryMockRecorder struct {
	mock *MockPortRepository
}

// NewMockPortRepository creates a new mock instance
func NewMockPortRepository(ctrl *gomock.Controller) *MockPortRepository {
	mock := &MockPortRepository{ctrl: ctrl}
	mock.recorder = &MockPortRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortRepository) EXPECT() *MockPortRepositoryMockRecorder {
	return m.recorder
}

// StorePorts mocks base method
func (m *MockPortRepository) StorePorts(arg0 []entity.PortInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorePorts", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorePorts indicates an expected call of StorePorts
func (mr *MockPortRepositoryMockRecorder) StorePorts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorePorts", reflect.TypeOf((*MockPortRepository)(nil).StorePorts), arg0)
}

// GetPorts mocks base method
func (m *MockPortRepository) GetPorts() ([]entity.PortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPorts")
	ret0, _ := ret[0].([]entity.PortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPorts indicates an expected call of GetPorts
func (mr *MockPortRepositoryMockRecorder) GetPorts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPorts", reflect.TypeOf((*MockPortRepository)(nil).GetPorts))
}
