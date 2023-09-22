// Code generated by MockGen. DO NOT EDIT.
// Source: ./user.go

// Package mockcontroller is a generated GoMock package.
package mockcontroller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockUserController is a mock of UserController interface.
type MockUserController struct {
	ctrl     *gomock.Controller
	recorder *MockUserControllerMockRecorder
}

// MockUserControllerMockRecorder is the mock recorder for MockUserController.
type MockUserControllerMockRecorder struct {
	mock *MockUserController
}

// NewMockUserController creates a new mock instance.
func NewMockUserController(ctrl *gomock.Controller) *MockUserController {
	mock := &MockUserController{ctrl: ctrl}
	mock.recorder = &MockUserControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserController) EXPECT() *MockUserControllerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserController) Get(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockUserControllerMockRecorder) Get(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserController)(nil).Get), c)
}

// GetEvents mocks base method.
func (m *MockUserController) GetEvents(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockUserControllerMockRecorder) GetEvents(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockUserController)(nil).GetEvents), c)
}

// GetExternalCalendars mocks base method.
func (m *MockUserController) GetExternalCalendars(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalCalendars", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetExternalCalendars indicates an expected call of GetExternalCalendars.
func (mr *MockUserControllerMockRecorder) GetExternalCalendars(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalCalendars", reflect.TypeOf((*MockUserController)(nil).GetExternalCalendars), c)
}
