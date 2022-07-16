// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/grimerssy/go-example/pkg/auth"

	core "github.com/grimerssy/go-example/internal/core"

	mock "github.com/stretchr/testify/mock"
)

// AuthUseCase is an autogenerated mock type for the authUseCase type
type AuthUseCase struct {
	mock.Mock
}

type AuthUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthUseCase) EXPECT() *AuthUseCase_Expecter {
	return &AuthUseCase_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: ctx, input
func (_m *AuthUseCase) Login(ctx context.Context, input *core.User) (auth.Tokens, error) {
	ret := _m.Called(ctx, input)

	var r0 auth.Tokens
	if rf, ok := ret.Get(0).(func(context.Context, *core.User) auth.Tokens); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.Tokens)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.User) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthUseCase_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AuthUseCase_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//  - ctx context.Context
//  - input *core.User
func (_e *AuthUseCase_Expecter) Login(ctx interface{}, input interface{}) *AuthUseCase_Login_Call {
	return &AuthUseCase_Login_Call{Call: _e.mock.On("Login", ctx, input)}
}

func (_c *AuthUseCase_Login_Call) Run(run func(ctx context.Context, input *core.User)) *AuthUseCase_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*core.User))
	})
	return _c
}

func (_c *AuthUseCase_Login_Call) Return(_a0 auth.Tokens, _a1 error) *AuthUseCase_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Signup provides a mock function with given fields: ctx, user
func (_m *AuthUseCase) Signup(ctx context.Context, user *core.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthUseCase_Signup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Signup'
type AuthUseCase_Signup_Call struct {
	*mock.Call
}

// Signup is a helper method to define mock.On call
//  - ctx context.Context
//  - user *core.User
func (_e *AuthUseCase_Expecter) Signup(ctx interface{}, user interface{}) *AuthUseCase_Signup_Call {
	return &AuthUseCase_Signup_Call{Call: _e.mock.On("Signup", ctx, user)}
}

func (_c *AuthUseCase_Signup_Call) Run(run func(ctx context.Context, user *core.User)) *AuthUseCase_Signup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*core.User))
	})
	return _c
}

func (_c *AuthUseCase_Signup_Call) Return(_a0 error) *AuthUseCase_Signup_Call {
	_c.Call.Return(_a0)
	return _c
}

type NewAuthUseCaseT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUseCase creates a new instance of AuthUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUseCase(t NewAuthUseCaseT) *AuthUseCase {
	mock := &AuthUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
