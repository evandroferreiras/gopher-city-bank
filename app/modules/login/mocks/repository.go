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

// GetAccountByCpf provides a mock function with given fields: cpf
func (_m *Repository) GetAccountByCpf(cpf string) (*model.Account, error) {
	ret := _m.Called(cpf)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(string) *model.Account); ok {
		r0 = rf(cpf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cpf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
