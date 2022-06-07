// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: accountId, operationTypeId, amount
func (_m *UseCase) Create(accountId int, operationTypeId int, amount float64) error {
	ret := _m.Called(accountId, operationTypeId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, float64) error); ok {
		r0 = rf(accountId, operationTypeId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewUseCaseT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t NewUseCaseT) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}