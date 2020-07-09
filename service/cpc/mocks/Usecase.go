// Code generated by mockery v1.0.0
package mocks

import (
	context "context"

	"github.com/models"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) GetAllCity(ctx context.Context, page, limit, offset int) (*models.CityDtoWithPagination, error) {
	ret := _m.Called(ctx, page, limit, offset)

	var r0 *models.CityDtoWithPagination
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) *models.CityDtoWithPagination); ok {
		r0 = rf(ctx, page, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CityDtoWithPagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) error); ok {
		r1 = rf(ctx, page, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetCityById(ctx context.Context, id int) (*models.CityDto, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.CityDto
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.CityDto); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CityDto)
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

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Usecase) CreateCity(ctx context.Context, f *models.NewCommandCity, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, f, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandCity, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, f, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandCity, string) error); ok {
		r1 = rf(ctx, f, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) UpdateCity(ctx context.Context, f *models.NewCommandCity, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, f, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandCity, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, f, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandCity, string) error); ok {
		r1 = rf(ctx, f, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) DeleteCity(ctx context.Context, id int, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, id, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, int, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, id, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, id, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) GetAllProvince(ctx context.Context, page, limit, offset int) (*models.ProvinceDtoWithPagination, error) {
	ret := _m.Called(ctx, page, limit, offset)

	var r0 *models.ProvinceDtoWithPagination
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) *models.ProvinceDtoWithPagination); ok {
		r0 = rf(ctx, page, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProvinceDtoWithPagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) error); ok {
		r1 = rf(ctx, page, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetProvinceById(ctx context.Context, id int) (*models.ProvinceDto, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.ProvinceDto
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.ProvinceDto); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProvinceDto)
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

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Usecase) CreateProvince(ctx context.Context, f *models.NewCommandProvince, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, f, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandProvince, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, f, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandProvince, string) error); ok {
		r1 = rf(ctx, f, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) UpdateProvince(ctx context.Context, f *models.NewCommandProvince, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, f, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandProvince, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, f, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandProvince, string) error); ok {
		r1 = rf(ctx, f, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) DeleteProvince(ctx context.Context, id int, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, id, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, int, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, id, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, id, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) GetAllCountry(ctx context.Context, page, limit, offset int) (*models.CountryDtoWithPagination, error) {
	ret := _m.Called(ctx, page, limit, offset)

	var r0 *models.CountryDtoWithPagination
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) *models.CountryDtoWithPagination); ok {
		r0 = rf(ctx, page, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CountryDtoWithPagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) error); ok {
		r1 = rf(ctx, page, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetCountryById(ctx context.Context, id int) (*models.CountryDto, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.CountryDto
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.CountryDto); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CountryDto)
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

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Usecase) CreateCountry(ctx context.Context, f *models.NewCommandCountry, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, f, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandCountry, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, f, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandCountry, string) error); ok {
		r1 = rf(ctx, f, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) UpdateCountry(ctx context.Context, f *models.NewCommandCountry, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, f, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandCountry, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, f, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandCountry, string) error); ok {
		r1 = rf(ctx, f, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) DeleteCountry(ctx context.Context, id int, token string) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, id, token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, int, string) *models.ResponseDelete); ok {
		r0 = rf(ctx, id, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, id, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
