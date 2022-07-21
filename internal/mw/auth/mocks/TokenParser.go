// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/grimerssy/go-example/pkg/auth"

	mock "github.com/stretchr/testify/mock"
)

// TokenParser is an autogenerated mock type for the TokenParser type
type TokenParser struct {
	mock.Mock
}

type TokenParser_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenParser) EXPECT() *TokenParser_Expecter {
	return &TokenParser_Expecter{mock: &_m.Mock}
}

// GetUserId provides a mock function with given fields: ctx, tokens
func (_m *TokenParser) GetUserId(ctx context.Context, tokens auth.AccessToken) (int64, error) {
	ret := _m.Called(ctx, tokens)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, auth.AccessToken) int64); ok {
		r0 = rf(ctx, tokens)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, auth.AccessToken) error); ok {
		r1 = rf(ctx, tokens)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenParser_GetUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserId'
type TokenParser_GetUserId_Call struct {
	*mock.Call
}

// GetUserId is a helper method to define mock.On call
//  - ctx context.Context
//  - tokens auth.AccessToken
func (_e *TokenParser_Expecter) GetUserId(ctx interface{}, tokens interface{}) *TokenParser_GetUserId_Call {
	return &TokenParser_GetUserId_Call{Call: _e.mock.On("GetUserId", ctx, tokens)}
}

func (_c *TokenParser_GetUserId_Call) Run(run func(ctx context.Context, tokens auth.AccessToken)) *TokenParser_GetUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(auth.AccessToken))
	})
	return _c
}

func (_c *TokenParser_GetUserId_Call) Return(_a0 int64, _a1 error) *TokenParser_GetUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewTokenParserT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenParser creates a new instance of TokenParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenParser(t NewTokenParserT) *TokenParser {
	mock := &TokenParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}