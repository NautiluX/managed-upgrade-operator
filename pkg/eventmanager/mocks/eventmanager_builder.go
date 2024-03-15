// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/eventmanager (interfaces: EventManagerBuilder)
//
// Generated by this command:
//
//	mockgen -destination=mocks/eventmanager_builder.go -package=mocks github.com/openshift/managed-upgrade-operator/pkg/eventmanager EventManagerBuilder
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	eventmanager "github.com/openshift/managed-upgrade-operator/pkg/eventmanager"
	gomock "go.uber.org/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockEventManagerBuilder is a mock of EventManagerBuilder interface.
type MockEventManagerBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockEventManagerBuilderMockRecorder
}

// MockEventManagerBuilderMockRecorder is the mock recorder for MockEventManagerBuilder.
type MockEventManagerBuilderMockRecorder struct {
	mock *MockEventManagerBuilder
}

// NewMockEventManagerBuilder creates a new mock instance.
func NewMockEventManagerBuilder(ctrl *gomock.Controller) *MockEventManagerBuilder {
	mock := &MockEventManagerBuilder{ctrl: ctrl}
	mock.recorder = &MockEventManagerBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventManagerBuilder) EXPECT() *MockEventManagerBuilderMockRecorder {
	return m.recorder
}

// NewManager mocks base method.
func (m *MockEventManagerBuilder) NewManager(arg0 client.Client) (eventmanager.EventManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewManager", arg0)
	ret0, _ := ret[0].(eventmanager.EventManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewManager indicates an expected call of NewManager.
func (mr *MockEventManagerBuilderMockRecorder) NewManager(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewManager", reflect.TypeOf((*MockEventManagerBuilder)(nil).NewManager), arg0)
}
