// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Exporter is an autogenerated mock type for the Exporter type
type Exporter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Exporter) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDefaultHeaders provides a mock function with given fields: w
func (_m *Exporter) SetDefaultHeaders(w http.ResponseWriter) {
	_m.Called(w)
}

// SetWriter provides a mock function with given fields: w
func (_m *Exporter) SetWriter(w http.ResponseWriter) {
	_m.Called(w)
}

// Write provides a mock function with given fields: _a0
func (_m *Exporter) Write(_a0 []string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewExporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewExporter creates a new instance of Exporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExporter(t mockConstructorTestingTNewExporter) *Exporter {
	mock := &Exporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
