// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	rates "github.com/ferruvich/go-exchange-rates-api/internal/rates"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepositorer is a mock of Repositorer interface
type MockRepositorer struct {
	ctrl     *gomock.Controller
	recorder *MockRepositorerMockRecorder
}

// MockRepositorerMockRecorder is the mock recorder for MockRepositorer
type MockRepositorerMockRecorder struct {
	mock *MockRepositorer
}

// NewMockRepositorer creates a new mock instance
func NewMockRepositorer(ctrl *gomock.Controller) *MockRepositorer {
	mock := &MockRepositorer{ctrl: ctrl}
	mock.recorder = &MockRepositorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositorer) EXPECT() *MockRepositorerMockRecorder {
	return m.recorder
}

// CurrentRates mocks base method
func (m *MockRepositorer) CurrentRates(base string) (*rates.BasedRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentRates", base)
	ret0, _ := ret[0].(*rates.BasedRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentRates indicates an expected call of CurrentRates
func (mr *MockRepositorerMockRecorder) CurrentRates(base interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentRates", reflect.TypeOf((*MockRepositorer)(nil).CurrentRates), base)
}

// CurrentSpecificRates mocks base method
func (m *MockRepositorer) CurrentSpecificRates(base string, currencies []string) (*rates.BasedRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentSpecificRates", base, currencies)
	ret0, _ := ret[0].(*rates.BasedRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentSpecificRates indicates an expected call of CurrentSpecificRates
func (mr *MockRepositorerMockRecorder) CurrentSpecificRates(base, currencies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentSpecificRates", reflect.TypeOf((*MockRepositorer)(nil).CurrentSpecificRates), base, currencies)
}

// HistoricalSpecificRates mocks base method
func (m *MockRepositorer) HistoricalSpecificRates(base, start, end string, currencies []string) (*rates.HistoricalRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HistoricalSpecificRates", base, start, end, currencies)
	ret0, _ := ret[0].(*rates.HistoricalRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistoricalSpecificRates indicates an expected call of HistoricalSpecificRates
func (mr *MockRepositorerMockRecorder) HistoricalSpecificRates(base, start, end, currencies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistoricalSpecificRates", reflect.TypeOf((*MockRepositorer)(nil).HistoricalSpecificRates), base, start, end, currencies)
}
