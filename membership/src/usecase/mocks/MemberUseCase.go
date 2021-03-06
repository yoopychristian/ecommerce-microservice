// Code generated by mockery v1.0.0
package mocks

import (
	model "ecommerce-microservice/membership/src/model"

	mock "github.com/stretchr/testify/mock"

	usecase "ecommerce-microservice/membership/src/usecase"
)

// MemberUseCase is an autogenerated mock type for the MemberUseCase type
type MemberUseCase struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *MemberUseCase) FindByEmail(email string) <-chan usecase.UseCaseResult {
	ret := _m.Called(email)

	var r0 <-chan usecase.UseCaseResult
	if rf, ok := ret.Get(0).(func(string) <-chan usecase.UseCaseResult); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan usecase.UseCaseResult)
		}
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *MemberUseCase) FindByID(id string) <-chan usecase.UseCaseResult {
	ret := _m.Called(id)

	var r0 <-chan usecase.UseCaseResult
	if rf, ok := ret.Get(0).(func(string) <-chan usecase.UseCaseResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan usecase.UseCaseResult)
		}
	}

	return r0
}

// Save provides a mock function with given fields: m
func (_m *MemberUseCase) Save(m *model.Member) <-chan error {
	ret := _m.Called(m)

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func(*model.Member) <-chan error); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}
