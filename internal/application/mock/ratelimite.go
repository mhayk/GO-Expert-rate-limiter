// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// RateLimiterServiceInterface is an autogenerated mock type for the RateLimiterServiceInterface type
type RateLimiterServiceInterface struct {
	mock.Mock
}

type RateLimiterServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RateLimiterServiceInterface) EXPECT() *RateLimiterServiceInterface_Expecter {
	return &RateLimiterServiceInterface_Expecter{mock: &_m.Mock}
}

// AllowRequest provides a mock function with given fields: ip, token
func (_m *RateLimiterServiceInterface) AllowRequest(ip string, token string) bool {
	ret := _m.Called(ip, token)

	if len(ret) == 0 {
		panic("no return value specified for AllowRequest")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(ip, token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RateLimiterServiceInterface_AllowRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllowRequest'
type RateLimiterServiceInterface_AllowRequest_Call struct {
	*mock.Call
}

// AllowRequest is a helper method to define mock.On call
//   - ip string
//   - token string
func (_e *RateLimiterServiceInterface_Expecter) AllowRequest(ip interface{}, token interface{}) *RateLimiterServiceInterface_AllowRequest_Call {
	return &RateLimiterServiceInterface_AllowRequest_Call{Call: _e.mock.On("AllowRequest", ip, token)}
}

func (_c *RateLimiterServiceInterface_AllowRequest_Call) Run(run func(ip string, token string)) *RateLimiterServiceInterface_AllowRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *RateLimiterServiceInterface_AllowRequest_Call) Return(_a0 bool) *RateLimiterServiceInterface_AllowRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RateLimiterServiceInterface_AllowRequest_Call) RunAndReturn(run func(string, string) bool) *RateLimiterServiceInterface_AllowRequest_Call {
	_c.Call.Return(run)
	return _c
}

// NewRateLimiterServiceInterface creates a new instance of RateLimiterServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRateLimiterServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RateLimiterServiceInterface {
	mock := &RateLimiterServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
