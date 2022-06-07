// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	pkg "github.com/anhaya/go-sample-api/pkg"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Atomic provides a mock function with given fields: fn
func (_m *Repository) Atomic(fn func(pkg.DBExecutor) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(pkg.DBExecutor) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
