// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// GreeterUseCase is an autogenerated mock type for the greeterUseCase type
type GreeterUseCase struct {
	mock.Mock
}

type GreeterUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *GreeterUseCase) EXPECT() *GreeterUseCase_Expecter {
	return &GreeterUseCase_Expecter{mock: &_m.Mock}
}

// Greet provides a mock function with given fields: ctx, userId
func (_m *GreeterUseCase) Greet(ctx context.Context, userId int64) (string, error) {
	ret := _m.Called(ctx, userId)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GreeterUseCase_Greet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Greet'
type GreeterUseCase_Greet_Call struct {
	*mock.Call
}

// Greet is a helper method to define mock.On call
//  - ctx context.Context
//  - userId int64
func (_e *GreeterUseCase_Expecter) Greet(ctx interface{}, userId interface{}) *GreeterUseCase_Greet_Call {
	return &GreeterUseCase_Greet_Call{Call: _e.mock.On("Greet", ctx, userId)}
}

func (_c *GreeterUseCase_Greet_Call) Run(run func(ctx context.Context, userId int64)) *GreeterUseCase_Greet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *GreeterUseCase_Greet_Call) Return(_a0 string, _a1 error) *GreeterUseCase_Greet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewGreeterUseCaseT interface {
	mock.TestingT
	Cleanup(func())
}

// NewGreeterUseCase creates a new instance of GreeterUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGreeterUseCase(t NewGreeterUseCaseT) *GreeterUseCase {
	mock := &GreeterUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}