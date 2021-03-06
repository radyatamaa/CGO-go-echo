package payment

import (
	"context"

	"github.com/models"
)

type Usecase interface {
	ConfirmPaymentBoarding(ctx context.Context,orderId string,token string)(*models.ResponseConfirmBoarding,error)
	Insert(ctx context.Context, payment *models.Transaction, token string, points float64,autoComplete bool) (string, error)
	ConfirmPayment(ctx context.Context, confirmIn *models.ConfirmPaymentIn) error
	ConfirmPaymentByDate(ctx context.Context,payment *models.ConfirmTransactionPayment)error
	SendingEmailConfirmPayment(ctx context.Context, confirmIn *models.ConfirmPaymentIn) error
	SendingEmailConfirmPaymentByDate(ctx context.Context,payment *models.ConfirmTransactionPayment)error
}
