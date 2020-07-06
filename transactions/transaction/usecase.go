package transaction

import (
	"context"

	"github.com/models"
)

type Usecase interface {
	//GetTransactionByDate(ctx context.Context,date string)()
	CountSuccess(ctx context.Context) (*models.Count, error)
	List(ctx context.Context, startDate, endDate, search, status string, page, limit, offset *int,token string,isAdmin bool,isTransportation bool,isExperience bool,isSchedule bool,tripType,paymentType,activityType string,confirmType string) (*models.TransactionWithPagination, error)
	CountThisMonth(ctx context.Context) (*models.TotalTransaction, error)
}
