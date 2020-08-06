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
// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) Fetch(ctx context.Context, cursor string, num int64) (res []*models.Experience, nextCursor string, err error) {
	ret := _m.Called(ctx, cursor,num)

	var r0 []*models.Experience
	if rf, ok := ret.Get(0).(func(context.Context, string,int64) []*models.Experience); ok {
		r0 = rf(ctx, cursor,num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Experience)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string,int64) string); ok {
		r1 = rf(ctx, cursor,num)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string,int64) error); ok {
		r2 = rf(ctx, cursor,num)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Error(12)
		}
	}
	return r0, r1,r2
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) SearchExp(ctx context.Context, harborID, cityID string) ([]*models.ExpSearch, error) {
	ret := _m.Called(ctx, harborID,cityID)

	var r0 []*models.ExpSearch
	if rf, ok := ret.Get(0).(func(context.Context, string,string) []*models.ExpSearch); ok {
		r0 = rf(ctx, harborID,cityID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ExpSearch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, harborID,cityID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetByID(ctx context.Context, id string) (*models.ExperienceJoinForegnKey, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.ExperienceJoinForegnKey
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.ExperienceJoinForegnKey); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ExperienceJoinForegnKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Repository) GetAllExperience(ctx context.Context) ([]*models.Experience, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Experience
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Experience); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Experience)
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
func (_m *Repository) SelectIdGetByMerchantId(ctx context.Context, merchantId string,month string,year int,date string) ([]*string, error) {
	ret := _m.Called(ctx, merchantId,month,year,date)

	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context, string,string,int,string) []*string); ok {
		r0 = rf(ctx, merchantId,month,year,date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string,int,string) error); ok {
		r1 = rf(ctx, merchantId,month,year,date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetUserDiscoverPreference(ctx context.Context, page *int, size *int) ([]*models.ExpUserDiscoverPreference, error) {
	ret := _m.Called(ctx, page,size)

	var r0 []*models.ExpUserDiscoverPreference
	if rf, ok := ret.Get(0).(func(context.Context, *int,*int) []*models.ExpUserDiscoverPreference); ok {
		r0 = rf(ctx, page,size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ExpUserDiscoverPreference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int,*int) error); ok {
		r1 = rf(ctx, page,size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetUserDiscoverPreferenceByHarborsIdOrProvince(ctx context.Context, harborsId *string,provinceId *int) ([]*models.ExpUserDiscoverPreference, error) {
	ret := _m.Called(ctx, harborsId,provinceId)

	var r0 []*models.ExpUserDiscoverPreference
	if rf, ok := ret.Get(0).(func(context.Context, *string,*int) []*models.ExpUserDiscoverPreference); ok {
		r0 = rf(ctx, harborsId,provinceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ExpUserDiscoverPreference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string,*int) error); ok {
		r1 = rf(ctx, harborsId,provinceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetIdByHarborsId(ctx context.Context, harborsId string) ([]*string, error) {
	ret := _m.Called(ctx, harborsId)

	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context, string) []*string); ok {
		r0 = rf(ctx, harborsId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, harborsId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetIdByCityId(ctx context.Context, cityId string) ([]*string, error) {
	ret := _m.Called(ctx, cityId)

	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context, string) []*string); ok {
		r0 = rf(ctx, cityId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cityId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) QueryFilterSearch(ctx context.Context, query string, limit, offset int) ([]*models.ExpSearch, error) {
	ret := _m.Called(ctx, query,limit,offset)

	var r0 []*models.ExpSearch
	if rf, ok := ret.Get(0).(func(context.Context, string,int,int) []*models.ExpSearch); ok {
		r0 = rf(ctx, query,limit,offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ExpSearch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,int,int) error); ok {
		r1 = rf(ctx, query,limit,offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetByCategoryID(ctx context.Context, categoryId int) ([]*models.ExpSearch, error) {
	ret := _m.Called(ctx, categoryId)

	var r0 []*models.ExpSearch
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.ExpSearch); ok {
		r0 = rf(ctx, categoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ExpSearch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) UpdateRating(ctx context.Context,exp models.Experience)error {
	ret := _m.Called(ctx, exp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Experience) error); ok {
		r0 = rf(ctx, exp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) UpdateStatus(ctx context.Context, status int,id string,user string)error {
	ret := _m.Called(ctx, status,id,user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int,string,string) error); ok {
		r0 = rf(ctx, status,id,user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) Update(ctx context.Context, a *models.Experience) (*string, error) {
	ret := _m.Called(ctx, a)

	var r0 *string
	if rf, ok := ret.Get(0).(func(context.Context, *models.Experience) *string); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Experience) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) Insert(ctx context.Context, a *models.Experience) (*string, error) {
	ret := _m.Called(ctx, a)

	var r0 *string
	if rf, ok := ret.Get(0).(func(context.Context, *models.Experience) *string); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Experience) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) Delete(ctx context.Context, id string, deleted_by string) error {
	ret := _m.Called(ctx, id,deleted_by)

	var r0 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r0 = rf(ctx, id,deleted_by)
	} else {
		r0 = ret.Error(1)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetSuccessBookCount(ctx context.Context, merchantId string) (int, error) {
	ret := _m.Called(ctx, merchantId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, merchantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetPublishedExpCount(ctx context.Context, merchantId string) (int, error) {
	ret := _m.Called(ctx, merchantId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, merchantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetExpPendingTransactionCount(ctx context.Context, merchantId string) (int, error) {
	ret := _m.Called(ctx, merchantId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, merchantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetExpFailedTransactionCount(ctx context.Context, merchantId string) (int, error) {
	ret := _m.Called(ctx, merchantId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, merchantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) CountFilterSearch(ctx context.Context, query string) (int, error) {
	ret := _m.Called(ctx, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetExpCount(ctx context.Context, merchantId string) (int, error) {
	ret := _m.Called(ctx, merchantId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, merchantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetExperienceByBookingId( ctx context.Context,bookingId ,experiencePaymentId string)(*models.ExperienceWithExperiencePayment,error) {
	ret := _m.Called(ctx, bookingId,experiencePaymentId)

	var r0 *models.ExperienceWithExperiencePayment
	if rf, ok := ret.Get(0).(func(context.Context, string,string) *models.ExperienceWithExperiencePayment); ok {
		r0 = rf(ctx, bookingId,experiencePaymentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ExperienceWithExperiencePayment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, bookingId,experiencePaymentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
