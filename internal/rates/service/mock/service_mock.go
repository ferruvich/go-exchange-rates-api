// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service_mock is a generated GoMock package.
package service_mock

import (
	rates "github.com/ferruvich/go-exchange-rates-api/internal/rates"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

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

// CurrentGBPUSDRates mocks base method
func (m *MockServicer) CurrentGBPUSDRates() ([]*rates.BasedRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentGBPUSDRates")
	ret0, _ := ret[0].([]*rates.BasedRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentGBPUSDRates indicates an expected call of CurrentGBPUSDRates
func (mr *MockServicerMockRecorder) CurrentGBPUSDRates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentGBPUSDRates", reflect.TypeOf((*MockServicer)(nil).CurrentGBPUSDRates))
}

// CurrentEURRate mocks base method
func (m *MockServicer) CurrentEURRate(currency string) (*rates.BasedRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentEURRate", currency)
	ret0, _ := ret[0].(*rates.BasedRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentEURRate indicates an expected call of CurrentEURRate
func (mr *MockServicerMockRecorder) CurrentEURRate(currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentEURRate", reflect.TypeOf((*MockServicer)(nil).CurrentEURRate), currency)
}

// RecommendEURExchange mocks base method
func (m *MockServicer) RecommendEURExchange(currency string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecommendEURExchange", currency)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecommendEURExchange indicates an expected call of RecommendEURExchange
func (mr *MockServicerMockRecorder) RecommendEURExchange(currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecommendEURExchange", reflect.TypeOf((*MockServicer)(nil).RecommendEURExchange), currency)
}
