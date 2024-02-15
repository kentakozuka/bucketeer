// Code generated by MockGen. DO NOT EDIT.
// Source: auto_ops_rule.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAutoOpsRuleStorage is a mock of AutoOpsRuleStorage interface.
type MockAutoOpsRuleStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAutoOpsRuleStorageMockRecorder
}

// MockAutoOpsRuleStorageMockRecorder is the mock recorder for MockAutoOpsRuleStorage.
type MockAutoOpsRuleStorageMockRecorder struct {
	mock *MockAutoOpsRuleStorage
}

// NewMockAutoOpsRuleStorage creates a new mock instance.
func NewMockAutoOpsRuleStorage(ctrl *gomock.Controller) *MockAutoOpsRuleStorage {
	mock := &MockAutoOpsRuleStorage{ctrl: ctrl}
	mock.recorder = &MockAutoOpsRuleStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoOpsRuleStorage) EXPECT() *MockAutoOpsRuleStorageMockRecorder {
	return m.recorder
}

// CountNotTriggeredAutoOpsRules mocks base method.
func (m *MockAutoOpsRuleStorage) CountNotTriggeredAutoOpsRules(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountNotTriggeredAutoOpsRules", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountNotTriggeredAutoOpsRules indicates an expected call of CountNotTriggeredAutoOpsRules.
func (mr *MockAutoOpsRuleStorageMockRecorder) CountNotTriggeredAutoOpsRules(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountNotTriggeredAutoOpsRules", reflect.TypeOf((*MockAutoOpsRuleStorage)(nil).CountNotTriggeredAutoOpsRules), ctx)
}