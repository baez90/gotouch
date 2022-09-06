// Code generated by MockGen. DO NOT EDIT.
// Source: ./store.go

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetStoreValues mocks base method.
func (m *MockStore) GetStoreValues() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreValues")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// GetStoreValues indicates an expected call of GetStoreValues.
func (mr *MockStoreMockRecorder) GetStoreValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreValues", reflect.TypeOf((*MockStore)(nil).GetStoreValues))
}

// GetValue mocks base method.
func (m *MockStore) GetValue(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetValue indicates an expected call of GetValue.
func (mr *MockStoreMockRecorder) GetValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockStore)(nil).GetValue), key)
}

// SetValue mocks base method.
func (m *MockStore) SetValue(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValue", key, value)
}

// SetValue indicates an expected call of SetValue.
func (mr *MockStoreMockRecorder) SetValue(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValue", reflect.TypeOf((*MockStore)(nil).SetValue), key, value)
}

// StoreValues mocks base method.
func (m *MockStore) StoreValues(key map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StoreValues", key)
}

// StoreValues indicates an expected call of StoreValues.
func (mr *MockStoreMockRecorder) StoreValues(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreValues", reflect.TypeOf((*MockStore)(nil).StoreValues), key)
}
