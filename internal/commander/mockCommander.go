// Code generated by MockGen. DO NOT EDIT.
// Source: ./commander.go

// Package commander is a generated GoMock package.
package commander

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommander is a mock of Commander interface.
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander.
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance.
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// CompressDirectory mocks base method.
func (m *MockCommander) CompressDirectory(opts *PackageCommandOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompressDirectory", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompressDirectory indicates an expected call of CompressDirectory.
func (mr *MockCommanderMockRecorder) CompressDirectory(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompressDirectory", reflect.TypeOf((*MockCommander)(nil).CompressDirectory), opts)
}

// CreateNewProject mocks base method.
func (m *MockCommander) CreateNewProject(opts *CreateCommandOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewProject", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewProject indicates an expected call of CreateNewProject.
func (mr *MockCommanderMockRecorder) CreateNewProject(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewProject", reflect.TypeOf((*MockCommander)(nil).CreateNewProject), opts)
}

// ValidateYaml mocks base method.
func (m *MockCommander) ValidateYaml(opts *ValidateCommandOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateYaml", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateYaml indicates an expected call of ValidateYaml.
func (mr *MockCommanderMockRecorder) ValidateYaml(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateYaml", reflect.TypeOf((*MockCommander)(nil).ValidateYaml), opts)
}