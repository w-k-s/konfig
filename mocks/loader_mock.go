// Code generated by MockGen. DO NOT EDIT.
// Source: ./loader.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	konfig "github.com/w-k-s/konfig"
)

// MockLoader is a mock of Loader interface.
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader.
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance.
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockLoader) Load(arg0 konfig.Values) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockLoaderMockRecorder) Load(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockLoader)(nil).Load), arg0)
}

// MaxRetry mocks base method.
func (m *MockLoader) MaxRetry() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxRetry")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxRetry indicates an expected call of MaxRetry.
func (mr *MockLoaderMockRecorder) MaxRetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxRetry", reflect.TypeOf((*MockLoader)(nil).MaxRetry))
}

// Name mocks base method.
func (m *MockLoader) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLoaderMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLoader)(nil).Name))
}

// RetryDelay mocks base method.
func (m *MockLoader) RetryDelay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryDelay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetryDelay indicates an expected call of RetryDelay.
func (mr *MockLoaderMockRecorder) RetryDelay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryDelay", reflect.TypeOf((*MockLoader)(nil).RetryDelay))
}

// StopOnFailure mocks base method.
func (m *MockLoader) StopOnFailure() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopOnFailure")
	ret0, _ := ret[0].(bool)
	return ret0
}

// StopOnFailure indicates an expected call of StopOnFailure.
func (mr *MockLoaderMockRecorder) StopOnFailure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopOnFailure", reflect.TypeOf((*MockLoader)(nil).StopOnFailure))
}
