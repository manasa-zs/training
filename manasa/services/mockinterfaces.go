// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	models "manasa/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockCar is a mock of Car interface.
type MockCar struct {
	ctrl     *gomock.Controller
	recorder *MockCarMockRecorder
}

// MockCarMockRecorder is the mock recorder for MockCar.
type MockCarMockRecorder struct {
	mock *MockCar
}

// NewMockCar creates a new mock instance.
func NewMockCar(ctrl *gomock.Controller) *MockCar {
	mock := &MockCar{ctrl: ctrl}
	mock.recorder = &MockCarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCar) EXPECT() *MockCarMockRecorder {
	return m.recorder
}

// CarCreate mocks base method.
func (m *MockCar) CarCreate(arg0 context.Context, arg1 *models.Car) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CarCreate", arg0, arg1)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CarCreate indicates an expected call of CarCreate.
func (mr *MockCarMockRecorder) CarCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CarCreate", reflect.TypeOf((*MockCar)(nil).CarCreate), arg0, arg1)
}

// CarDelete mocks base method.
func (m *MockCar) CarDelete(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CarDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CarDelete indicates an expected call of CarDelete.
func (mr *MockCarMockRecorder) CarDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CarDelete", reflect.TypeOf((*MockCar)(nil).CarDelete), arg0, arg1)
}

// CarGetByBrand mocks base method.
func (m *MockCar) CarGetByBrand(arg0 context.Context, arg1 string) ([]*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CarGetByBrand", arg0, arg1)
	ret0, _ := ret[0].([]*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CarGetByBrand indicates an expected call of CarGetByBrand.
func (mr *MockCarMockRecorder) CarGetByBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CarGetByBrand", reflect.TypeOf((*MockCar)(nil).CarGetByBrand), arg0, arg1)
}

// CarGetByID mocks base method.
func (m *MockCar) CarGetByID(arg0 context.Context, arg1 uuid.UUID) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CarGetByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CarGetByID indicates an expected call of CarGetByID.
func (mr *MockCarMockRecorder) CarGetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CarGetByID", reflect.TypeOf((*MockCar)(nil).CarGetByID), arg0, arg1)
}

// CarUpdate mocks base method.
func (m *MockCar) CarUpdate(arg0 context.Context, arg1 *models.Car) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CarUpdate", arg0, arg1)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CarUpdate indicates an expected call of CarUpdate.
func (mr *MockCarMockRecorder) CarUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CarUpdate", reflect.TypeOf((*MockCar)(nil).CarUpdate), arg0, arg1)
}

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// EngineCreate mocks base method.
func (m *MockEngine) EngineCreate(arg0 context.Context, arg1 *models.Engine) (*models.Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EngineCreate", arg0, arg1)
	ret0, _ := ret[0].(*models.Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EngineCreate indicates an expected call of EngineCreate.
func (mr *MockEngineMockRecorder) EngineCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EngineCreate", reflect.TypeOf((*MockEngine)(nil).EngineCreate), arg0, arg1)
}

// EngineDelete mocks base method.
func (m *MockEngine) EngineDelete(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EngineDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EngineDelete indicates an expected call of EngineDelete.
func (mr *MockEngineMockRecorder) EngineDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EngineDelete", reflect.TypeOf((*MockEngine)(nil).EngineDelete), arg0, arg1)
}

// EngineGetByID mocks base method.
func (m *MockEngine) EngineGetByID(arg0 context.Context, arg1 uuid.UUID) (*models.Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EngineGetByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EngineGetByID indicates an expected call of EngineGetByID.
func (mr *MockEngineMockRecorder) EngineGetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EngineGetByID", reflect.TypeOf((*MockEngine)(nil).EngineGetByID), arg0, arg1)
}

// EngineUpdate mocks base method.
func (m *MockEngine) EngineUpdate(arg0 context.Context, arg1 *models.Engine) (*models.Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EngineUpdate", arg0, arg1)
	ret0, _ := ret[0].(*models.Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EngineUpdate indicates an expected call of EngineUpdate.
func (mr *MockEngineMockRecorder) EngineUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EngineUpdate", reflect.TypeOf((*MockEngine)(nil).EngineUpdate), arg0, arg1)
}
