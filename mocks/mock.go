package mocks

import "github.com/stretchr/testify/mock"

type ExampleService struct {
	mock.Mock
}

func (_m ExampleService) Get(no int) int {
	ret := _m.Called(no)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(no)
	} else {
		r0 = ret.Get(0).(int)
	}
	return r0
}
