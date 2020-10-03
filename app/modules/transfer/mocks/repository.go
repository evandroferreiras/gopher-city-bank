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

// CommitTransaction provides a mock function with given fields:
func (_m *Repository) CommitTransaction() {
	_m.Called()
}

// GetAccount provides a mock function with given fields: id
func (_m *Repository) GetAccount(id string) (model.Account, error) {
	ret := _m.Called(id)

	var r0 model.Account
	if rf, ok := ret.Get(0).(func(string) model.Account); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RollbackTransaction provides a mock function with given fields:
func (_m *Repository) RollbackTransaction() {
	_m.Called()
}

// StartTransaction provides a mock function with given fields:
func (_m *Repository) StartTransaction() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAccountBalance provides a mock function with given fields: id, newBalance
func (_m *Repository) UpdateAccountBalance(id string, newBalance float64) error {
	ret := _m.Called(id, newBalance)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64) error); ok {
		r0 = rf(id, newBalance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
