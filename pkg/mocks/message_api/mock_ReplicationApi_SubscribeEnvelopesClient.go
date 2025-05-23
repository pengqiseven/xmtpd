// Code generated by mockery v2.52.2. DO NOT EDIT.

package message_api

import (
	context "context"

	message_api "github.com/xmtp/xmtpd/pkg/proto/xmtpv4/message_api"
	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// MockReplicationApi_SubscribeEnvelopesClient is an autogenerated mock type for the ReplicationApi_SubscribeEnvelopesClient type
type MockReplicationApi_SubscribeEnvelopesClient struct {
	mock.Mock
}

type MockReplicationApi_SubscribeEnvelopesClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReplicationApi_SubscribeEnvelopesClient) EXPECT() *MockReplicationApi_SubscribeEnvelopesClient_Expecter {
	return &MockReplicationApi_SubscribeEnvelopesClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with no fields
func (_m *MockReplicationApi_SubscribeEnvelopesClient) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) CloseSend() *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call) Run(run func()) *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call) Return(_a0 error) *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call) RunAndReturn(run func() error) *MockReplicationApi_SubscribeEnvelopesClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with no fields
func (_m *MockReplicationApi_SubscribeEnvelopesClient) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockReplicationApi_SubscribeEnvelopesClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockReplicationApi_SubscribeEnvelopesClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) Context() *MockReplicationApi_SubscribeEnvelopesClient_Context_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Context_Call) Run(run func()) *MockReplicationApi_SubscribeEnvelopesClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Context_Call) Return(_a0 context.Context) *MockReplicationApi_SubscribeEnvelopesClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Context_Call) RunAndReturn(run func() context.Context) *MockReplicationApi_SubscribeEnvelopesClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with no fields
func (_m *MockReplicationApi_SubscribeEnvelopesClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReplicationApi_SubscribeEnvelopesClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockReplicationApi_SubscribeEnvelopesClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) Header() *MockReplicationApi_SubscribeEnvelopesClient_Header_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Header_Call) Run(run func()) *MockReplicationApi_SubscribeEnvelopesClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *MockReplicationApi_SubscribeEnvelopesClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *MockReplicationApi_SubscribeEnvelopesClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with no fields
func (_m *MockReplicationApi_SubscribeEnvelopesClient) Recv() (*message_api.SubscribeEnvelopesResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *message_api.SubscribeEnvelopesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*message_api.SubscribeEnvelopesResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *message_api.SubscribeEnvelopesResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message_api.SubscribeEnvelopesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReplicationApi_SubscribeEnvelopesClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockReplicationApi_SubscribeEnvelopesClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) Recv() *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call) Run(run func()) *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call) Return(_a0 *message_api.SubscribeEnvelopesResponse, _a1 error) *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call) RunAndReturn(run func() (*message_api.SubscribeEnvelopesResponse, error)) *MockReplicationApi_SubscribeEnvelopesClient_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockReplicationApi_SubscribeEnvelopesClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) RecvMsg(m interface{}) *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call) Run(run func(m interface{})) *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call) Return(_a0 error) *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockReplicationApi_SubscribeEnvelopesClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockReplicationApi_SubscribeEnvelopesClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) SendMsg(m interface{}) *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call) Run(run func(m interface{})) *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call) Return(_a0 error) *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockReplicationApi_SubscribeEnvelopesClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with no fields
func (_m *MockReplicationApi_SubscribeEnvelopesClient) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockReplicationApi_SubscribeEnvelopesClient_Expecter) Trailer() *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call {
	return &MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call) Run(run func()) *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call) Return(_a0 metadata.MD) *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *MockReplicationApi_SubscribeEnvelopesClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReplicationApi_SubscribeEnvelopesClient creates a new instance of MockReplicationApi_SubscribeEnvelopesClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReplicationApi_SubscribeEnvelopesClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReplicationApi_SubscribeEnvelopesClient {
	mock := &MockReplicationApi_SubscribeEnvelopesClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
