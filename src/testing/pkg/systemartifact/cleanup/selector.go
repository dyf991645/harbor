// Code generated by mockery v2.35.4. DO NOT EDIT.

package cleanup

import (
	context "context"

	model "github.com/goharbor/harbor/src/pkg/systemartifact/model"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Selector is an autogenerated mock type for the Selector type
type Selector struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx
func (_m *Selector) List(ctx context.Context) ([]*model.SystemArtifact, error) {
	ret := _m.Called(ctx)

	var r0 []*model.SystemArtifact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.SystemArtifact, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.SystemArtifact); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.SystemArtifact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWithFilters provides a mock function with given fields: ctx, query
func (_m *Selector) ListWithFilters(ctx context.Context, query *q.Query) ([]*model.SystemArtifact, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.SystemArtifact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*model.SystemArtifact, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.SystemArtifact); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.SystemArtifact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSelector creates a new instance of Selector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSelector(t interface {
	mock.TestingT
	Cleanup(func())
}) *Selector {
	mock := &Selector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
