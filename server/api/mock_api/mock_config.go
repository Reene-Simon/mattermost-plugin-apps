// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-apps/server/api (interfaces: Configurator)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/mattermost/mattermost-plugin-apps/server/api"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
)

// MockConfigurator is a mock of Configurator interface
type MockConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockConfiguratorMockRecorder
}

// MockConfiguratorMockRecorder is the mock recorder for MockConfigurator
type MockConfiguratorMockRecorder struct {
	mock *MockConfigurator
}

// NewMockConfigurator creates a new mock instance
func NewMockConfigurator(ctrl *gomock.Controller) *MockConfigurator {
	mock := &MockConfigurator{ctrl: ctrl}
	mock.recorder = &MockConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigurator) EXPECT() *MockConfiguratorMockRecorder {
	return m.recorder
}

// GetConfig mocks base method
func (m *MockConfigurator) GetConfig() api.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(api.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockConfiguratorMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockConfigurator)(nil).GetConfig))
}

// GetMattermostConfig mocks base method
func (m *MockConfigurator) GetMattermostConfig() *model.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMattermostConfig")
	ret0, _ := ret[0].(*model.Config)
	return ret0
}

// GetMattermostConfig indicates an expected call of GetMattermostConfig
func (mr *MockConfiguratorMockRecorder) GetMattermostConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMattermostConfig", reflect.TypeOf((*MockConfigurator)(nil).GetMattermostConfig))
}

// Reconfigure mocks base method
func (m *MockConfigurator) Reconfigure(arg0 api.StoredConfig, arg1 ...api.Configurable) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Reconfigure", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconfigure indicates an expected call of Reconfigure
func (mr *MockConfiguratorMockRecorder) Reconfigure(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconfigure", reflect.TypeOf((*MockConfigurator)(nil).Reconfigure), varargs...)
}

// StoreConfig mocks base method
func (m *MockConfigurator) StoreConfig(arg0 api.StoredConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreConfig indicates an expected call of StoreConfig
func (mr *MockConfiguratorMockRecorder) StoreConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreConfig", reflect.TypeOf((*MockConfigurator)(nil).StoreConfig), arg0)
}