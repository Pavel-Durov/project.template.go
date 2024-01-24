// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sqlc "p3ld3v.dev/template/app/services/db/sqlc"
)

// DbStore is an autogenerated mock type for the DbStore type
type DbStore struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *DbStore) Connect() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: name
func (_m *DbStore) CreateUser(name string) (*sqlc.User, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *sqlc.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*sqlc.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *sqlc.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlc.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: id
func (_m *DbStore) GetUser(id int64) (*sqlc.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *sqlc.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*sqlc.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *sqlc.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlc.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDbStore creates a new instance of DbStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDbStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *DbStore {
	mock := &DbStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
