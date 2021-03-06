// Code generated by mockery v1.0.0
package mocks

import (
	model "ecommerce-microservice/membership/src/model"

	mock "github.com/stretchr/testify/mock"

	repository "ecommerce-microservice/membership/src/repository"
)

// MembershipRepository is an autogenerated mock type for the MembershipRepository type
type MembershipRepository struct {
	mock.Mock
}

// Load provides a mock function with given fields: id
func (_m *MembershipRepository) Load(id string) <-chan repository.RepositoryResult {
	ret := _m.Called(id)

	var r0 <-chan repository.RepositoryResult
	if rf, ok := ret.Get(0).(func(string) <-chan repository.RepositoryResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan repository.RepositoryResult)
		}
	}

	return r0
}

// Save provides a mock function with given fields: m
func (_m *MembershipRepository) Save(m *model.Member) <-chan error {
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
