// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch-daemon/internal/backend (interfaces: NerdctlNetworkSvc)

// Package mocks_backend is a generated GoMock package.
package mocks_backend

import (
	context "context"
	reflect "reflect"

	types "github.com/containerd/nerdctl/v2/pkg/api/types"
	dockercompat "github.com/containerd/nerdctl/v2/pkg/inspecttypes/dockercompat"
	netutil "github.com/containerd/nerdctl/v2/pkg/netutil"
	libcni "github.com/containernetworking/cni/libcni"
	types0 "github.com/containernetworking/cni/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockNerdctlNetworkSvc is a mock of NerdctlNetworkSvc interface.
type MockNerdctlNetworkSvc struct {
	ctrl     *gomock.Controller
	recorder *MockNerdctlNetworkSvcMockRecorder
}

// MockNerdctlNetworkSvcMockRecorder is the mock recorder for MockNerdctlNetworkSvc.
type MockNerdctlNetworkSvcMockRecorder struct {
	mock *MockNerdctlNetworkSvc
}

// NewMockNerdctlNetworkSvc creates a new mock instance.
func NewMockNerdctlNetworkSvc(ctrl *gomock.Controller) *MockNerdctlNetworkSvc {
	mock := &MockNerdctlNetworkSvc{ctrl: ctrl}
	mock.recorder = &MockNerdctlNetworkSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNerdctlNetworkSvc) EXPECT() *MockNerdctlNetworkSvcMockRecorder {
	return m.recorder
}

// AddNetworkList mocks base method.
func (m *MockNerdctlNetworkSvc) AddNetworkList(arg0 context.Context, arg1 *libcni.NetworkConfigList, arg2 *libcni.RuntimeConf) (types0.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNetworkList", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNetworkList indicates an expected call of AddNetworkList.
func (mr *MockNerdctlNetworkSvcMockRecorder) AddNetworkList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetworkList", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).AddNetworkList), arg0, arg1, arg2)
}

// CreateNetwork mocks base method.
func (m *MockNerdctlNetworkSvc) CreateNetwork(arg0 types.NetworkCreateOptions) (*netutil.NetworkConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetwork", arg0)
	ret0, _ := ret[0].(*netutil.NetworkConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetwork indicates an expected call of CreateNetwork.
func (mr *MockNerdctlNetworkSvcMockRecorder) CreateNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetwork", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).CreateNetwork), arg0)
}

// FilterNetworks mocks base method.
func (m *MockNerdctlNetworkSvc) FilterNetworks(arg0 func(*netutil.NetworkConfig) bool) ([]*netutil.NetworkConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterNetworks", arg0)
	ret0, _ := ret[0].([]*netutil.NetworkConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterNetworks indicates an expected call of FilterNetworks.
func (mr *MockNerdctlNetworkSvcMockRecorder) FilterNetworks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterNetworks", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).FilterNetworks), arg0)
}

// InspectNetwork mocks base method.
func (m *MockNerdctlNetworkSvc) InspectNetwork(arg0 context.Context, arg1 *netutil.NetworkConfig) (*dockercompat.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectNetwork", arg0, arg1)
	ret0, _ := ret[0].(*dockercompat.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectNetwork indicates an expected call of InspectNetwork.
func (mr *MockNerdctlNetworkSvcMockRecorder) InspectNetwork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectNetwork", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).InspectNetwork), arg0, arg1)
}

// Namespace mocks base method.
func (m *MockNerdctlNetworkSvc) Namespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockNerdctlNetworkSvcMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).Namespace))
}

// NetconfPath mocks base method.
func (m *MockNerdctlNetworkSvc) NetconfPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetconfPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// NetconfPath indicates an expected call of NetconfPath.
func (mr *MockNerdctlNetworkSvcMockRecorder) NetconfPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetconfPath", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).NetconfPath))
}

// RemoveNetwork mocks base method.
func (m *MockNerdctlNetworkSvc) RemoveNetwork(arg0 *netutil.NetworkConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNetwork indicates an expected call of RemoveNetwork.
func (mr *MockNerdctlNetworkSvcMockRecorder) RemoveNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNetwork", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).RemoveNetwork), arg0)
}

// UsedNetworkInfo mocks base method.
func (m *MockNerdctlNetworkSvc) UsedNetworkInfo(arg0 context.Context) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsedNetworkInfo", arg0)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsedNetworkInfo indicates an expected call of UsedNetworkInfo.
func (mr *MockNerdctlNetworkSvcMockRecorder) UsedNetworkInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsedNetworkInfo", reflect.TypeOf((*MockNerdctlNetworkSvc)(nil).UsedNetworkInfo), arg0)
}
