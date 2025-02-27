// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	dynamic "k8s.io/client-go/dynamic"

	mock "github.com/stretchr/testify/mock"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// KubeClient is an autogenerated mock type for the KubeClient type
type KubeClient struct {
	mock.Mock
}

// ClusterID provides a mock function with given fields:
func (_m *KubeClient) ClusterID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resource provides a mock function with given fields: resource
func (_m *KubeClient) Resource(resource schema.GroupVersionResource) dynamic.NamespaceableResourceInterface {
	ret := _m.Called(resource)

	var r0 dynamic.NamespaceableResourceInterface
	if rf, ok := ret.Get(0).(func(schema.GroupVersionResource) dynamic.NamespaceableResourceInterface); ok {
		r0 = rf(resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.NamespaceableResourceInterface)
		}
	}

	return r0
}

type mockConstructorTestingTNewKubeClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewKubeClient creates a new instance of KubeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKubeClient(t mockConstructorTestingTNewKubeClient) *KubeClient {
	mock := &KubeClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
