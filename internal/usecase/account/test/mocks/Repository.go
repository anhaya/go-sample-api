// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entity "github.com/anhaya/go-sample-api/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: documentNumber, balance
func (_m *Repository) Create(documentNumber string, balance float64) (int64, error) {
	ret := _m.Called(documentNumber, balance)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, float64) int64); ok {
		r0 = rf(documentNumber, balance)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, float64) error); ok {
		r1 = rf(documentNumber, balance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: accountId
func (_m *Repository) Get(accountId int) (entity.Account, error) {
	ret := _m.Called(accountId)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(int) entity.Account); ok {
		r0 = rf(accountId)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: accountId, balance
func (_m *Repository) Update(accountId int, balance float64) error {
	ret := _m.Called(accountId, balance)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, float64) error); ok {
		r0 = rf(accountId, balance)
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