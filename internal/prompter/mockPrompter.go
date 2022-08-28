// Code generated by MockGen. DO NOT EDIT.
// Source: ./prompter.go

// Package prompter is a generated GoMock package.
package prompter

import (
	"fmt"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPrompter is a mock of Prompter interface.
type MockPrompter struct {
	ctrl     *gomock.Controller
	recorder *MockPrompterMockRecorder
}

// MockPrompterMockRecorder is the mock recorder for MockPrompter.
type MockPrompterMockRecorder struct {
	mock *MockPrompter
}

// NewMockPrompter creates a new mock instance.
func NewMockPrompter(ctrl *gomock.Controller) *MockPrompter {
	mock := &MockPrompter{ctrl: ctrl}
	mock.recorder = &MockPrompterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrompter) EXPECT() *MockPrompterMockRecorder {
	return m.recorder
}

// AskForSelectionFromList mocks base method.
func (m *MockPrompter) AskForSelectionFromList(direction string, list []fmt.Stringer) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskForSelectionFromList", direction, list)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskForSelectionFromList indicates an expected call of AskForSelectionFromList.
func (mr *MockPrompterMockRecorder) AskForSelectionFromList(direction, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskForSelectionFromList", reflect.TypeOf((*MockPrompter)(nil).AskForSelectionFromList), direction, list)
}

// AskForString mocks base method.
func (m *MockPrompter) AskForString(direction string, validator StringValidator) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskForString", direction, validator)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskForString indicates an expected call of AskForString.
func (mr *MockPrompterMockRecorder) AskForString(direction, validator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskForString", reflect.TypeOf((*MockPrompter)(nil).AskForString), direction, validator)
}

// MockOption is a mock of Choice interface.
type MockOption struct {
	ctrl     *gomock.Controller
	recorder *MockOptionMockRecorder
}

// MockOptionMockRecorder is the mock recorder for MockOption.
type MockOptionMockRecorder struct {
	mock *MockOption
}

// NewMockOption creates a new mock instance.
func NewMockOption(ctrl *gomock.Controller) *MockOption {
	mock := &MockOption{ctrl: ctrl}
	mock.recorder = &MockOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOption) EXPECT() *MockOptionMockRecorder {
	return m.recorder
}

// String mocks base method.
func (m *MockOption) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockOptionMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockOption)(nil).String))
}
