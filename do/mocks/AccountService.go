package mocks

import "github.com/bryanl/doit/do"
import "github.com/stretchr/testify/mock"

type AccountService struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *AccountService) Get() (*do.Account, error) {
	ret := _m.Called()

	var r0 *do.Account
	if rf, ok := ret.Get(0).(func() *do.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RateLimit provides a mock function with given fields:
func (_m *AccountService) RateLimit() (*do.RateLimit, error) {
	ret := _m.Called()

	var r0 *do.RateLimit
	if rf, ok := ret.Get(0).(func() *do.RateLimit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.RateLimit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
