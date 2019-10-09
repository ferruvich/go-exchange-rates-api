// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service_mock is a generated GoMock package.
package service_mock

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	http "net/http"
	reflect "reflect"
)

// MockDoer is a mock of Doer interface
type MockDoer struct {
	ctrl     *gomock.Controller
	recorder *MockDoerMockRecorder
}

// MockDoerMockRecorder is the mock recorder for MockDoer
type MockDoerMockRecorder struct {
	mock *MockDoer
}

// NewMockDoer creates a new mock instance
func NewMockDoer(ctrl *gomock.Controller) *MockDoer {
	mock := &MockDoer{ctrl: ctrl}
	mock.recorder = &MockDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer) EXPECT() *MockDoerMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockDoer) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockDoerMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockDoer)(nil).Do), req)
}

// MockServicer is a mock of Servicer interface
type MockServicer struct {
	ctrl     *gomock.Controller
	recorder *MockServicerMockRecorder
}

// MockServicerMockRecorder is the mock recorder for MockServicer
type MockServicerMockRecorder struct {
	mock *MockServicer
}

// NewMockServicer creates a new mock instance
func NewMockServicer(ctrl *gomock.Controller) *MockServicer {
	mock := &MockServicer{ctrl: ctrl}
	mock.recorder = &MockServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServicer) EXPECT() *MockServicerMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockServicer) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockServicerMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockServicer)(nil).Do), req)
}

// NewRequest mocks base method
func (m *MockServicer) NewRequest(method, url string, body io.ReadCloser, qp map[string]string) (*http.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRequest", method, url, body, qp)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRequest indicates an expected call of NewRequest
func (mr *MockServicerMockRecorder) NewRequest(method, url, body, qp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRequest", reflect.TypeOf((*MockServicer)(nil).NewRequest), method, url, body, qp)
}
