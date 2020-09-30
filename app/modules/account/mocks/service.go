// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	model "github.com/evandroferreiras/gopher-city-bank/app/model"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Service) Create(_a0 *model.NewAccount) (*model.Account, error) {
	ret := _m.Called(_a0)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(*model.NewAccount) *model.Account); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.NewAccount) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}