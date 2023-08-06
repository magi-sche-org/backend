// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth.go

// Package mockmiddleware is a generated GoMock package.
package mockmiddleware

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockAuthMiddleware is a mock of AuthMiddleware interface.
type MockAuthMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMiddlewareMockRecorder
}

// MockAuthMiddlewareMockRecorder is the mock recorder for MockAuthMiddleware.
type MockAuthMiddlewareMockRecorder struct {
	mock *MockAuthMiddleware
}

// NewMockAuthMiddleware creates a new mock instance.
func NewMockAuthMiddleware(ctrl *gomock.Controller) *MockAuthMiddleware {
	mock := &MockAuthMiddleware{ctrl: ctrl}
	mock.recorder = &MockAuthMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthMiddleware) EXPECT() *MockAuthMiddlewareMockRecorder {
	return m.recorder
}

// Handler mocks base method.
func (m *MockAuthMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler", next)
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// Handler indicates an expected call of Handler.
func (mr *MockAuthMiddlewareMockRecorder) Handler(next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockAuthMiddleware)(nil).Handler), next)
}