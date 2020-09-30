// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	model "github.com/evandroferreiras/gopher-city-bank/app/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: newAccount
func (_m *Repository) Create(newAccount model.NewAccount) (*model.AccountCreated, error) {
	ret := _m.Called(newAccount)

	var r0 *model.AccountCreated
	if rf, ok := ret.Get(0).(func(model.NewAccount) *model.AccountCreated); ok {
		r0 = rf(newAccount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccountCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.NewAccount) error); ok {
		r1 = rf(newAccount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
