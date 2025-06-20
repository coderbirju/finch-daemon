// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch-daemon/internal/service/network/driver (interfaces: DriverHandler)
//
// Generated by this command:
//
//	mockgen --destination=../../../../mocks/mocks_network/driver.go -package=mocks_network -mock_names DriverHandler=DriverHandler . DriverHandler
//

// Package mocks_network is a generated GoMock package.
package mocks_network

import (
	reflect "reflect"

	types "github.com/containerd/nerdctl/v2/pkg/api/types"
	netutil "github.com/containerd/nerdctl/v2/pkg/netutil"
	types0 "github.com/runfinch/finch-daemon/api/types"
	gomock "go.uber.org/mock/gomock"
)

// DriverHandler is a mock of DriverHandler interface.
type DriverHandler struct {
	ctrl     *gomock.Controller
	recorder *DriverHandlerMockRecorder
	isgomock struct{}
}

// DriverHandlerMockRecorder is the mock recorder for DriverHandler.
type DriverHandlerMockRecorder struct {
	mock *DriverHandler
}

// NewDriverHandler creates a new mock instance.
func NewDriverHandler(ctrl *gomock.Controller) *DriverHandler {
	mock := &DriverHandler{ctrl: ctrl}
	mock.recorder = &DriverHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *DriverHandler) EXPECT() *DriverHandlerMockRecorder {
	return m.recorder
}

// HandleCreateOptions mocks base method.
func (m *DriverHandler) HandleCreateOptions(request types0.NetworkCreateRequest, options types.NetworkCreateOptions) (types.NetworkCreateOptions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCreateOptions", request, options)
	ret0, _ := ret[0].(types.NetworkCreateOptions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleCreateOptions indicates an expected call of HandleCreateOptions.
func (mr *DriverHandlerMockRecorder) HandleCreateOptions(request, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCreateOptions", reflect.TypeOf((*DriverHandler)(nil).HandleCreateOptions), request, options)
}

// HandlePostCreate mocks base method.
func (m *DriverHandler) HandlePostCreate(net *netutil.NetworkConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandlePostCreate", net)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandlePostCreate indicates an expected call of HandlePostCreate.
func (mr *DriverHandlerMockRecorder) HandlePostCreate(net any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlePostCreate", reflect.TypeOf((*DriverHandler)(nil).HandlePostCreate), net)
}

// HandleRemove mocks base method.
func (m *DriverHandler) HandleRemove(net *netutil.NetworkConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleRemove", net)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleRemove indicates an expected call of HandleRemove.
func (mr *DriverHandlerMockRecorder) HandleRemove(net any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRemove", reflect.TypeOf((*DriverHandler)(nil).HandleRemove), net)
}
