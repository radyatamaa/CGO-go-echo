package repository

import (
	"context"
	"database/sql"

	guuid "github.com/google/uuid"
	"github.com/models"
	"github.com/transactions/payment"
)

type paymentRepository struct {
	Conn *sql.DB
}

// NewPaymentRepository will create an object that represent the article.repository interface
func NewPaymentRepository(Conn *sql.DB) payment.Repository {
	return &paymentRepository{Conn}
}
func (p paymentRepository) ChangeStatusTransByDate(ctx context.Context, payment *models.ConfirmTransactionPayment) error {
	if payment.IsCancelExp == true {
		query4 := `
	UPDATE
		transactions,
		booking_exps
	SET
		transactions.status = 4,
		transactions.remarks = ?
	WHERE
		(booking_exp_id = booking_exps.id OR booking_exps.order_id = transactions.order_id) 
		AND transactions.status != 2
		AND DATE(booking_exps.booking_date) = ? `

		if payment.TransId != "" {
			query4 = query4 + ` AND booking_exps.trans_id = '` + payment.TransId + `' `
		} else if payment.ExpId != "" {
			query4 = query4 + ` AND booking_exps.exp_id = '` + payment.ExpId + `' `
		}

		query8 := `
	UPDATE
		transactions,
		booking_exps
	SET
		transactions.status = 8,
		transactions.remarks = ?
	WHERE
		(booking_exp_id = booking_exps.id OR booking_exps.order_id = transactions.order_id)
		AND transactions.status = 2
		AND DATE(booking_exps.booking_date) = ? `

		if payment.TransId != "" {
			query8 = query8 + ` AND booking_exps.trans_id = '` + payment.TransId + `' `
		} else if payment.ExpId != "" {
			query8 = query8 + ` AND booking_exps.exp_id = '` + payment.ExpId + `' `
		}

		stmt, err := p.Conn.PrepareContext(ctx, query4)
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx,
			payment.Remarks,
			payment.BookingDate,
		)
		if err != nil {
			return err
		}

		stmt, err = p.Conn.PrepareContext(ctx, query8)
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx,
			payment.Remarks,
			payment.BookingDate,
		)
		if err != nil {
			return err
		}
	}else {
		if payment.BookingStatus == 0 {
			query := `
	UPDATE
		transactions,
		booking_exps
	SET
		transactions.status = ?,
		transactions.remarks = ?
	WHERE
		(booking_exp_id = booking_exps.id OR booking_exps.order_id = transactions.order_id)
		AND DATE(booking_exps.booking_date) = ? `

			if payment.TransId != "" {
				query = query + ` AND booking_exps.trans_id = '` + payment.TransId + `'`
			} else if payment.ExpId != "" {
				query = query + ` AND booking_exps.exp_id = '` + payment.ExpId + `'`
			}
			stmt, err := p.Conn.PrepareContext(ctx, query)
			if err != nil {
				return err
			}

			_, err = stmt.ExecContext(ctx,
				payment.TransactionStatus,
				payment.Remarks,
				payment.BookingDate,
			)
			if err != nil {
				return err
			}

		}else {
			query := `
	UPDATE
		transactions,
		booking_exps
	SET
		transactions.status = ?,
		transactions.remarks = ?,
		booking_exps.status = ?
	WHERE
		(booking_exp_id = booking_exps.id OR booking_exps.order_id = transactions.order_id)
		AND DATE(booking_exps.booking_date) = ? `

			if payment.TransId != "" {
				query = query + ` AND booking_exps.trans_id = '` + payment.TransId + `'`
			} else if payment.ExpId != "" {
				query = query + ` AND booking_exps.exp_id = '` + payment.ExpId + `'`
			}
			stmt, err := p.Conn.PrepareContext(ctx, query)
			if err != nil {
				return err
			}

			_, err = stmt.ExecContext(ctx,
				payment.TransactionStatus,
				payment.Remarks,
				payment.BookingStatus,
				payment.BookingDate,
			)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (p paymentRepository) Insert(ctx context.Context, pay *models.Transaction) (*models.Transaction, error) {
	id := guuid.New()
	pay.Id = id.String()
	q := `INSERT transactions SET id = ?, created_by = ?, created_date = ?, modified_by = ?, 
	modified_date = ?, deleted_by = ?, deleted_date = ?, is_deleted = ?, is_active = ?, booking_type = ?, 
	booking_exp_id = ?, promo_id = ?, payment_method_id = ?, experience_payment_id = ?, status = ?, total_price = ?, 
	currency = ?, order_id = ?,ex_change_rates=?,ex_change_currency=?,points=?,original_price=?`

	stmt, err := p.Conn.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(
		ctx,
		pay.Id,
		pay.CreatedBy,
		pay.CreatedDate,
		pay.ModifiedBy,
		pay.ModifiedDate,
		pay.DeletedBy,
		pay.DeletedDate,
		pay.IsDeleted,
		pay.IsActive,
		pay.BookingType,
		pay.BookingExpId,
		pay.PromoId,
		pay.PaymentMethodId,
		pay.ExperiencePaymentId,
		pay.Status,
		pay.TotalPrice,
		pay.Currency,
		pay.OrderId,
		pay.ExChangeRates,
		pay.ExChangeCurrency,
		pay.Points,
		pay.OriginalPrice,
	)
	if err != nil {
		return nil, err
	}

	return pay, nil
}

func (p paymentRepository) ConfirmPayment(ctx context.Context, confirmIn *models.ConfirmPaymentIn) error {
	if confirmIn.Amount != nil {
		query := `
	UPDATE
		transactions,
		booking_exps
	SET
		transactions.status = ?,
		transactions.total_price = ?,
		transactions.remarks = ?,
		booking_exps.status = ?
	WHERE
		(booking_exp_id = booking_exps.id OR booking_exps.order_id = transactions.order_id)
		AND transactions.id = ?`

		stmt, err := p.Conn.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx,
			confirmIn.TransactionStatus,
			confirmIn.Amount,
			confirmIn.Remarks,
			confirmIn.BookingStatus,
			confirmIn.TransactionID,
		)
		if err != nil {
			return err
		}

	} else {
		query := `
	UPDATE
		transactions,
		booking_exps
	SET
		transactions.status = ?,
		transactions.remarks = ?,
		booking_exps.status = ?
	WHERE
		(booking_exp_id = booking_exps.id OR booking_exps.order_id = transactions.order_id)
		AND transactions.id = ?`

		stmt, err := p.Conn.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx,
			confirmIn.TransactionStatus,
			confirmIn.Remarks,
			confirmIn.BookingStatus,
			confirmIn.TransactionID,
		)
		if err != nil {
			return err
		}

	}

	return nil
}
