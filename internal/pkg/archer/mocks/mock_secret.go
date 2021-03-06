// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/archer/secret.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSecretsManager is a mock of SecretsManager interface
type MockSecretsManager struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerMockRecorder
}

// MockSecretsManagerMockRecorder is the mock recorder for MockSecretsManager
type MockSecretsManagerMockRecorder struct {
	mock *MockSecretsManager
}

// NewMockSecretsManager creates a new mock instance
func NewMockSecretsManager(ctrl *gomock.Controller) *MockSecretsManager {
	mock := &MockSecretsManager{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretsManager) EXPECT() *MockSecretsManagerMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method
func (m *MockSecretsManager) CreateSecret(secretName, secretString string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", secretName, secretString)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockSecretsManagerMockRecorder) CreateSecret(secretName, secretString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretsManager)(nil).CreateSecret), secretName, secretString)
}

// DeleteSecret mocks base method
func (m *MockSecretsManager) DeleteSecret(secretName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", secretName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockSecretsManagerMockRecorder) DeleteSecret(secretName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretsManager)(nil).DeleteSecret), secretName)
}

// MockSecretCreator is a mock of SecretCreator interface
type MockSecretCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSecretCreatorMockRecorder
}

// MockSecretCreatorMockRecorder is the mock recorder for MockSecretCreator
type MockSecretCreatorMockRecorder struct {
	mock *MockSecretCreator
}

// NewMockSecretCreator creates a new mock instance
func NewMockSecretCreator(ctrl *gomock.Controller) *MockSecretCreator {
	mock := &MockSecretCreator{ctrl: ctrl}
	mock.recorder = &MockSecretCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretCreator) EXPECT() *MockSecretCreatorMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method
func (m *MockSecretCreator) CreateSecret(secretName, secretString string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", secretName, secretString)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockSecretCreatorMockRecorder) CreateSecret(secretName, secretString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretCreator)(nil).CreateSecret), secretName, secretString)
}

// MockSecretDeleter is a mock of SecretDeleter interface
type MockSecretDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSecretDeleterMockRecorder
}

// MockSecretDeleterMockRecorder is the mock recorder for MockSecretDeleter
type MockSecretDeleterMockRecorder struct {
	mock *MockSecretDeleter
}

// NewMockSecretDeleter creates a new mock instance
func NewMockSecretDeleter(ctrl *gomock.Controller) *MockSecretDeleter {
	mock := &MockSecretDeleter{ctrl: ctrl}
	mock.recorder = &MockSecretDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretDeleter) EXPECT() *MockSecretDeleterMockRecorder {
	return m.recorder
}

// DeleteSecret mocks base method
func (m *MockSecretDeleter) DeleteSecret(secretName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", secretName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockSecretDeleterMockRecorder) DeleteSecret(secretName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretDeleter)(nil).DeleteSecret), secretName)
}
