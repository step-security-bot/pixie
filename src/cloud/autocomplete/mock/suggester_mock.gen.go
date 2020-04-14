// Code generated by MockGen. DO NOT EDIT.
// Source: autocomplete.go

// Package mock_autocomplete is a generated GoMock package.
package mock_autocomplete

import (
	gomock "github.com/golang/mock/gomock"
	autocomplete "pixielabs.ai/pixielabs/src/cloud/autocomplete"
	reflect "reflect"
)

// MockSuggester is a mock of Suggester interface
type MockSuggester struct {
	ctrl     *gomock.Controller
	recorder *MockSuggesterMockRecorder
}

// MockSuggesterMockRecorder is the mock recorder for MockSuggester
type MockSuggesterMockRecorder struct {
	mock *MockSuggester
}

// NewMockSuggester creates a new mock instance
func NewMockSuggester(ctrl *gomock.Controller) *MockSuggester {
	mock := &MockSuggester{ctrl: ctrl}
	mock.recorder = &MockSuggesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSuggester) EXPECT() *MockSuggesterMockRecorder {
	return m.recorder
}

// GetSuggestions mocks base method
func (m *MockSuggester) GetSuggestions(reqs []*autocomplete.SuggestionRequest) ([]*autocomplete.SuggestionResult, error) {
	ret := m.ctrl.Call(m, "GetSuggestions", reqs)
	ret0, _ := ret[0].([]*autocomplete.SuggestionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuggestions indicates an expected call of GetSuggestions
func (mr *MockSuggesterMockRecorder) GetSuggestions(reqs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestions", reflect.TypeOf((*MockSuggester)(nil).GetSuggestions), reqs)
}
