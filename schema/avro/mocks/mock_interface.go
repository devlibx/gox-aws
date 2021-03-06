// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_gox_aws_avro is a generated GoMock package.
package mock_gox_aws_avro

import (
	reflect "reflect"

	gox "github.com/devlibx/gox-base"
	gomock "github.com/golang/mock/gomock"
)

// MockSchemaEngine is a mock of SchemaEngine interface.
type MockSchemaEngine struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaEngineMockRecorder
}

// MockSchemaEngineMockRecorder is the mock recorder for MockSchemaEngine.
type MockSchemaEngineMockRecorder struct {
	mock *MockSchemaEngine
}

// NewMockSchemaEngine creates a new mock instance.
func NewMockSchemaEngine(ctrl *gomock.Controller) *MockSchemaEngine {
	mock := &MockSchemaEngine{ctrl: ctrl}
	mock.recorder = &MockSchemaEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaEngine) EXPECT() *MockSchemaEngineMockRecorder {
	return m.recorder
}

// FromAvro mocks base method.
func (m *MockSchemaEngine) FromAvro(data []byte) (gox.StringObjectMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromAvro", data)
	ret0, _ := ret[0].(gox.StringObjectMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromAvro indicates an expected call of FromAvro.
func (mr *MockSchemaEngineMockRecorder) FromAvro(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromAvro", reflect.TypeOf((*MockSchemaEngine)(nil).FromAvro), data)
}

// ToAvro mocks base method.
func (m *MockSchemaEngine) ToAvro(data interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToAvro", data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToAvro indicates an expected call of ToAvro.
func (mr *MockSchemaEngineMockRecorder) ToAvro(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToAvro", reflect.TypeOf((*MockSchemaEngine)(nil).ToAvro), data)
}
