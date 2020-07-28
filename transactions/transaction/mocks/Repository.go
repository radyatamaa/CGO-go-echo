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

func (_m *Repository) GetIdTransactionByStatus(ctx context.Context, transactionStatus int) ([]*string, error) {
	ret := _m.Called(ctx,transactionStatus)

	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context, int) []*string); ok {
		r0 = rf(ctx,transactionStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx,transactionStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetTransactionByExpIdORTransId(ctx context.Context,date string,expId string,transId string,merchantId string,status string)([]*models.TransactionOut, error) {
	ret := _m.Called(ctx,date,expId,transId,merchantId,status)

	var r0 []*models.TransactionOut
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) []*models.TransactionOut); ok {
		r0 = rf(ctx,date,expId,transId,merchantId,status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TransactionOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string) error); ok {
		r1 = rf(ctx,date,expId,transId,merchantId,status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) GetTransactionByDate(ctx context.Context,date string ,isExperience bool,isTransportation bool,merchantId string)([]*models.TransactionByDate,error) {
	ret := _m.Called(ctx, date, isExperience,isTransportation,merchantId)

	var r0 []*models.TransactionByDate
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, bool, string) []*models.TransactionByDate); ok {
		r0 = rf(ctx, date, isExperience,isTransportation,merchantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TransactionByDate)
		}
	}


	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context,string ,bool,bool,string) error); ok {
		r1 = rf(ctx, date, isExperience,isTransportation,merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0,r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetCountTransactionByPromoId(ctx context.Context,promoId string,userId string)(int,error) {
	ret := _m.Called(ctx, promoId, userId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int); ok {
		r0 = rf(ctx, promoId,userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, promoId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetTransactionDownPaymentByDate(ctx context.Context)([]*models.TransactionWithBooking,error) {
	ret := _m.Called(ctx)

	var r0 []*models.TransactionWithBooking
	if rf, ok := ret.Get(0).(func(context.Context) []*models.TransactionWithBooking); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TransactionWithBooking)
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

// Store provides a mock function with given fields: ctx, a
func (_m *Repository) GetIdTransactionExpired(ctx context.Context)([]*string ,error) {
	ret := _m.Called(ctx)

	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context) []*string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
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

// Update provides a mock function with given fields: ctx, ar
func (_m *Repository) GetCountByExpId(ctx context.Context, date string, expId string,isTransaction bool) ([]*string, error) {
	ret := _m.Called(ctx, date,expId,isTransaction)

	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context) []*string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0,r1
}

func (_m *Repository) GetCountByTransId(ctx context.Context, transId string,isTransaction bool,date string) ([]*string, error) {
	ret := _m.Called(ctx,transId,isTransaction,date)


	var r0 []*string
	if rf, ok := ret.Get(0).(func(context.Context,string,bool,string) []*string); ok {
		r0 = rf(ctx,transId,isTransaction,date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context,string,bool,string) error); ok {
		r1 = rf(ctx,transId,isTransaction,date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id string) (*models.TransactionWMerchant, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.TransactionWMerchant
	if rf, ok := ret.Get(0).(func(context.Context,string) *models.TransactionWMerchant); ok {
		r0 = rf(ctx,id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TransactionWMerchant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context,string) error); ok {
		r1 = rf(ctx,id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) GetByBookingDate(ctx context.Context, bookingDate string,transId string,expId string) ([]*models.TransactionWMerchant, error) {
	ret := _m.Called(ctx, bookingDate, transId,expId)

	var r0 []*models.TransactionWMerchant
	if rf, ok := ret.Get(0).(func(context.Context, string, string,string) []*models.TransactionWMerchant); ok {
		r0 = rf(ctx, bookingDate,transId,expId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TransactionWMerchant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string,string) error); ok {
		r1 = rf(ctx, bookingDate,transId,expId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) CountSuccess(ctx context.Context) (int, error) {
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

// Store provides a mock function with given fields: ctx, a
func (_m *Repository) Count(ctx context.Context, startDate, endDate, search, status string, merchantId string,isTransportation bool,isExperience bool,isSchedule bool,tripType,paymentType,activityType string,confirmType string,class string,departureTimeStart string,departureTimeEnd string,arrivalTimeStart string,arrivalTimeEnd string,transactionId string) (int, error) {
	ret := _m.Called(ctx, startDate, endDate, search, status , merchantId ,isTransportation ,isExperience ,isSchedule ,tripType,paymentType,activityType ,confirmType ,class,departureTimeStart,departureTimeEnd,arrivalTimeStart,arrivalTimeEnd,transactionId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string,string,bool,bool,bool,string,string,string,string,string, string,string,string,string,string) int); ok {
		r0 = rf(ctx, startDate, endDate, search, status , merchantId ,isTransportation ,isExperience ,isSchedule ,tripType,paymentType,activityType ,confirmType ,class,departureTimeStart,departureTimeEnd,arrivalTimeStart,arrivalTimeEnd,transactionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string,string,bool,bool,bool,string,string,string,string,string, string,string,string,string,string) error); ok {
		r1 = rf(ctx, startDate, endDate, search, status , merchantId ,isTransportation ,isExperience ,isSchedule ,tripType,paymentType,activityType ,confirmType ,class,departureTimeStart,departureTimeEnd,arrivalTimeStart,arrivalTimeEnd,transactionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ar
func (_m *Repository) List(ctx context.Context, startDate, endDate, search, status string, limit, offset *int, merchantId string,isTransportation bool,isExperience bool,isSchedule bool,tripType,paymentType,activityType string,confirmType string,class string,departureTimeStart string,departureTimeEnd string,arrivalTimeStart string,arrivalTimeEnd string,transactionId string) ([]*models.TransactionOut, error) {
	ret := _m.Called(ctx, startDate, endDate, search, status , limit, offset , merchantId ,isTransportation,isExperience,isSchedule,tripType,paymentType,activityType ,confirmType ,class ,departureTimeStart,departureTimeEnd ,arrivalTimeStart ,arrivalTimeEnd,transactionId)

	var r0 []*models.TransactionOut
	if rf, ok := ret.Get(0).(func(context.Context,  string, string, string, string, *int,*int, string,bool,bool,bool,string,string,string, string,string,string,string,string,string,string) []*models.TransactionOut); ok {
		r0 = rf(ctx, startDate, endDate, search, status , limit, offset , merchantId ,isTransportation,isExperience,isSchedule,tripType,paymentType,activityType ,confirmType ,class ,departureTimeStart,departureTimeEnd ,arrivalTimeStart ,arrivalTimeEnd,transactionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TransactionOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context,  string, string, string, string, *int,*int, string,bool,bool,bool,string,string,string, string,string,string,string,string,string,string) error); ok {
		r1 = rf(ctx, startDate, endDate, search, status , limit, offset , merchantId ,isTransportation,isExperience,isSchedule,tripType,paymentType,activityType ,confirmType ,class ,departureTimeStart,departureTimeEnd ,arrivalTimeStart ,arrivalTimeEnd,transactionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) CountThisMonth(ctx context.Context) (*models.TotalTransaction, error) {
	ret := _m.Called(ctx)

	var r0 *models.TotalTransaction
	if rf, ok := ret.Get(0).(func(context.Context) *models.TotalTransaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TotalTransaction)
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
func (_m *Repository) UpdateAfterPayment(ctx context.Context, status int, vaNumber string, transactionId, bookingId string) error {
	ret := _m.Called(ctx,  status , vaNumber, transactionId, bookingId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string,string,string) error); ok {
		r0 = rf(ctx,  status , vaNumber, transactionId, bookingId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
