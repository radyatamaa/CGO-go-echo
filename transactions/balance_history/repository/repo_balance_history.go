package repository

import (
	"database/sql"
	guuid "github.com/google/uuid"
	"github.com/models"
	"github.com/sirupsen/logrus"
	"github.com/transactions/balance_history"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type balanceHistoryRepository struct {
	Conn *sql.DB
}



// NewpromoRepository will create an object that represent the article.Repository interface
func NewbalanceHistoryRepository(Conn *sql.DB) balance_history.Repository {
	return &balanceHistoryRepository{Conn}
}
func (m *balanceHistoryRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.BalanceHistory, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.BalanceHistory, 0)
	for rows.Next() {
		t := new(models.BalanceHistory)
		err = rows.Scan(
			&t.Id,
			&t.CreatedBy,
			&t.CreatedDate,
			&t.ModifiedBy,
			&t.ModifiedDate,
			&t.DeletedBy,
			&t.DeletedDate,
			&t.IsDeleted,
			&t.IsActive,
			&t.Status  ,
			&t.MerchantId  ,
			&t.Amount      ,
			&t.DateOfRequest ,
			&t.DateOfPayment ,
			&t.Remarks       ,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
func (b balanceHistoryRepository) GetAll(ctx context.Context, merchantId string,status string, limit, offset *int) ([]*models.BalanceHistory, error) {
	query := `SELECT * FROM balance_histories WHERE is_deleted = 0 AND is_active = 1`
	if merchantId != ""{
		query = query + ` AND merchant_id = '` + merchantId + `'`
	}
	if status != ""{
		query = query + ` AND status = '` + status + `'`
	}
	if limit != nil && offset != nil{
		query = query + ` LIMIT ` + strconv.Itoa(*offset) + ` , ` + strconv.Itoa(*limit)
	}
	res, err := b.fetch(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	return res, nil
}
func (m *balanceHistoryRepository) Count(ctx context.Context, merchantId string, status string) (int, error) {
	query := `SELECT COUNT(*) as count FROM balance_histories WHERE is_deleted = 0 AND is_active = 1`
	if merchantId != ""{
		query = query + ` AND merchant_id = '` + merchantId + `'`
	}
	if status != ""{
		query = query + ` AND status = '` + status + `'`
	}
	//if limit != nil && offset != nil{
	//	query = query + ` LIMIT ` + string(*limit) + ` , ` + string(*offset)
	//}
	rows, err := m.Conn.QueryContext(ctx, query)
	count, err := checkCount(rows)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return count, nil
}
func (b balanceHistoryRepository) Insert(ctx context.Context, a models.BalanceHistory) (*string, error) {
	a.Id = guuid.New().String()
	query := `INSERT balance_histories SET id=? , created_by=? , created_date=? , modified_by=?, modified_date=? , deleted_by=? , 
				deleted_date=? , is_deleted=? , is_active=? , status=?,merchant_id=?,amount=?,
				date_of_request=?,date_of_payment=?,remarks=?`
	stmt, err := b.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(ctx, a.Id, a.CreatedBy, time.Now(), nil, nil, nil, nil, 0, 1,a.Status,a.MerchantId,a.Amount,
							a.DateOfRequest,a.DateOfPayment,a.Remarks)
	if err != nil {
		return nil, err
	}

	//lastID, err := res.RowsAffected()
	//if err != nil {
	//	return err
	//}

	//a.Id = lastID
	return &a.Id, nil
}
func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}