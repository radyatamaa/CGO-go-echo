// Code generated by mockery v1.0.0
package mocks

import (
	context "context"

	"github.com/models"

	mock "github.com/stretchr/testify/mock"
)

// repository is an autogenerated mock type for the repository type
type Repository struct {
	mock.Mock
}

func (_m *Repository) GetCount(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int, deletedBy string) error {
	ret := _m.Called(ctx, id, deletedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, id, deletedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Repository) List(ctx context.Context) ([]*models.Rule, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Rule
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Rule); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) Fetch(ctx context.Context, limit int, offset int) ([]*models.Rule, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []*models.Rule
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*models.Rule); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id int) (*models.Rule, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Rule
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Rule); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *Repository) Insert(ctx context.Context, a *models.Rule) (*int, error) {
	ret := _m.Called(ctx, a)

	var r0 *int
	if rf, ok := ret.Get(0).(func(context.Context, *models.Rule) *int); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Rule) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ar
func (_m *Repository) Update(ctx context.Context, ar *models.Rule) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Rule) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
