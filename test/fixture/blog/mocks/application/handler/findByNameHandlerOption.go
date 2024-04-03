// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// findByNameHandlerOption is an autogenerated mock type for the findByNameHandlerOption type
type findByNameHandlerOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *findByNameHandlerOption) Execute(_a0 *handler.findByNameHandler) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*handler.findByNameHandler) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newFindByNameHandlerOption creates a new instance of findByNameHandlerOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newFindByNameHandlerOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *findByNameHandlerOption {
	mock := &findByNameHandlerOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}