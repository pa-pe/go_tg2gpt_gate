// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// ICache is an autogenerated mock type for the ICache type
type ICache struct {
	mock.Mock
}

// ClearNamespace provides a mock function with given fields: ctx, namespace
func (_m *ICache) ClearNamespace(ctx context.Context, namespace string) {
	_m.Called(ctx, namespace)
}

// Delete provides a mock function with given fields: ctx, namespace, key
func (_m *ICache) Delete(ctx context.Context, namespace string, key string) {
	_m.Called(ctx, namespace, key)
}

// Load provides a mock function with given fields: ctx, namespace, key, dest
func (_m *ICache) Load(ctx context.Context, namespace string, key string, dest interface{}) error {
	ret := _m.Called(ctx, namespace, key, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, namespace, key, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadList provides a mock function with given fields: ctx, namespace, key, dest, total
func (_m *ICache) LoadList(ctx context.Context, namespace string, key string, dest interface{}, total *int64) error {
	ret := _m.Called(ctx, namespace, key, dest, total)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}, *int64) error); ok {
		r0 = rf(ctx, namespace, key, dest, total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: ctx, namespace, key, value, period
func (_m *ICache) Put(ctx context.Context, namespace string, key string, value interface{}, period time.Duration) {
	_m.Called(ctx, namespace, key, value, period)
}

// PutList provides a mock function with given fields: ctx, namespace, key, value, total, period
func (_m *ICache) PutList(ctx context.Context, namespace string, key string, value interface{}, total int64, period time.Duration) {
	_m.Called(ctx, namespace, key, value, total, period)
}

type mockConstructorTestingTNewICache interface {
	mock.TestingT
	Cleanup(func())
}

// NewICache creates a new instance of ICache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICache(t mockConstructorTestingTNewICache) *ICache {
	mock := &ICache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
