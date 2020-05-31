package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/models"

	"github.com/sirupsen/logrus"
	"github.com/transactions/transaction"
)

type transactionRepository struct {
	Conn *sql.DB
}

func NewTransactionRepository(Conn *sql.DB) transaction.Repository {
	return &transactionRepository{Conn: Conn}
}

func (t transactionRepository) GetTransactionDownPaymentByDate(ctx context.Context) ([]*models.TransactionWithBooking, error) {
	threeDay := time.Now().AddDate(0, 0, 3).Format("2006-01-02")
	twoDay := time.Now().AddDate(0, 0, 2).Format("2006-01-02")
	tenDay := time.Now().AddDate(0, 0, 10).Format("2006-01-02")
	threetyDay := time.Now().AddDate(0, 0, 30).Format("2006-01-02")

	query := `
	SELECT 
		e.exp_title,
		be.booked_by,
		be.booked_by_email,
		be.booking_date,
		t.total_price,
		ep.price ,
		e.exp_duration,
		t.order_id,
		m.merchant_name,
		m.phone_number as merchant_phone
	FROM transactions t
	JOIN experience_payments ep on ep.id = t.experience_payment_id
	JOIN booking_exps be on be.id = t.booking_exp_id
	JOIN experiences e on e.id = be.exp_id
	JOIN merchants m on m.id = e.merchant_id	
	WHERE 
		ep.exp_payment_type_id = '86e71b8d-acc3-4ade-80c0-de67b9100633' AND 
		t.total_price != ep.price AND 
		(DATE(be.booking_date) = ? OR DATE(be.booking_date) = ? OR DATE(be.booking_date) = ? OR DATE(be.booking_date) = ?)`

	rows, err := t.Conn.QueryContext(ctx, query, threeDay, twoDay, tenDay, threetyDay)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result := make([]*models.TransactionWithBooking, 0)
	for rows.Next() {
		t := new(models.TransactionWithBooking)
		err = rows.Scan(
			&t.ExpTitle,
			&t.BookedBy,
			&t.BookedByEmail,
			&t.BookingDate,
			&t.TotalPrice,
			&t.Price,
			&t.ExpDuration,
			&t.OrderId,
			&t.MerchantName,
			&t.MerchantPhone,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
func (t transactionRepository) GetCountByExpId(ctx context.Context, date string, expId string) (*string, error) {
	query := `
	select b.guest_desc from transactions a
										join booking_exps b on a.order_id = b.order_id 
										join experiences c on c.id = b.exp_id 
										where a.status < 3 and (b.status = 1 or b.status = 3)
										and date(b.booking_date) = ? and exp_id = ?`

	rows, err := t.Conn.QueryContext(ctx, query, date, expId)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var bookingDesc string
	for rows.Next() {
		err = rows.Scan(&bookingDesc)
		if err != nil {
			return nil, err
		}
	}

	return &bookingDesc, nil
}
func (t transactionRepository) GetCountByTransId(ctx context.Context, transId string) (int, error) {
	query := `select count(*) from transactions a
											join booking_exps b on a.order_id = b.order_id 
											join transportations c on c.id = b.trans_id 
											join schedules d on b.schedule_id = d.id
											where a.status < 3 and b.trans_id = ?`

	rows, err := t.Conn.QueryContext(ctx, query, transId)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	count, err := checkCount(rows)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return count, nil
}
func (t transactionRepository) UpdateAfterPayment(ctx context.Context, status int, vaNumber string, transactionId, bookingId string) error {
	query := `UPDATE transactions SET status = ?, va_number = ? WHERE (id = ? OR booking_exp_id = ? OR order_id = ?)`

	stmt, err := t.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, status, vaNumber, transactionId, bookingId, bookingId)
	if err != nil {
		return err
	}

	return nil
}

func (t transactionRepository) CountThisMonth(ctx context.Context) (*models.TotalTransaction, error) {
	query := `
	SELECT
		count(CAST(created_date AS DATE)) as transaction_count,
		SUM(total_price) as transaction_value_total
	FROM
		transactions
	WHERE
		is_deleted = 0
		AND is_active = 1
		AND status = 2
		AND created_date BETWEEN date_add(CURRENT_DATE, interval - DAY(CURRENT_DATE) + 1 DAY)
		AND CURRENT_DATE`

	rows, err := t.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	total := new(models.TotalTransaction)
	for rows.Next() {
		err = rows.Scan(&total.TransactionCount, &total.TransactionValueTotal)
		if err != nil {
			return nil, err
		}
	}

	return total, nil
}

func (t transactionRepository) List(ctx context.Context, startDate, endDate, search, status string, limit, offset *int, merchantId string, isTransportation bool, isExperience bool, isSchedule bool) ([]*models.TransactionOut, error) {
	var transactionStatus int
	var bookingStatus int

	query := `
	SELECT
		t.id as transaction_id,
		e.id as exp_id,
		e.exp_type,
		e.exp_title,
		booking_exp_id,
		b.order_id as booking_code,
		b.created_date as booking_date,
		b.booking_date as check_in_date,
		b.booked_by,
		guest_desc,
		b.booked_by_email as email,
		t.status as transaction_status,
		b.status as booking_status,
		t.total_price,
		t.experience_payment_id as experience_payment_id,
		merchant_name,
		b.order_id,
		e.exp_duration,
		p.province_name,
		co.country_name
	FROM
		transactions t
		JOIN booking_exps b ON t.booking_exp_id = b.id
		JOIN experiences e ON b.exp_id = e.id
		JOIN merchants m ON e.merchant_id = m.id
		JOIN harbors  h ON e.harbors_id = h.id
		JOIN cities  c ON h.city_id = c.id
		JOIN provinces p on c.province_id = p.id
		JOIN countries co on p.country_id = co.id
	WHERE 
		t.is_deleted = 0
		AND t.is_active = 1
	`

	queryT := `
	SELECT
		t.id AS transaction_id,
		trans_id,
		trans_name,
		trans_title,
		booking_exp_id,
		b.order_id AS booking_code,
		b.created_date AS booking_date,
		b.booking_date AS check_in_date,
		b.booked_by,
		guest_desc,
		b.booked_by_email AS email,
		t.status AS transaction_status,
		b.status AS booking_status,
		t.total_price,
		tr.class as trans_class,
		merchant_name,
		b.order_id,
		tr.trans_capacity as exp_duration,
		trans_name as province_name,
		trans_title as country_name
	FROM
		transactions t
		JOIN booking_exps b ON t.booking_exp_id = b.id
		JOIN transportations tr ON b.trans_id = tr.id
		JOIN merchants m ON tr.merchant_id = m.id
	WHERE
		t.is_deleted = 0
		AND t.is_active = 1`

	if merchantId != "" {
		query = query + ` AND e.merchant_id = '` + merchantId + `' `
		queryT = queryT + ` AND tr.merchant_id = '` + merchantId + `' `
	}

	if search != "" {
		keyword := `'%` + search + `%'`
		query = query + ` AND (LOWER(b.booked_by) LIKE LOWER(` + keyword + `) OR LOWER(b.order_id) LIKE LOWER(` + keyword + `))`
		queryT = queryT + ` AND (LOWER(b.booked_by) LIKE LOWER(` + keyword + `) OR LOWER(b.order_id) LIKE LOWER(` + keyword + `))`
	}
	if startDate != "" && endDate != "" {
		if isSchedule == true {
			query = query + ` AND DATE(b.booking_date) BETWEEN '` + startDate + `' AND '` + endDate + `'`
			queryT = queryT + ` AND DATE(b.booking_date) BETWEEN '` + startDate + `' AND '` + endDate + `'`
		} else {
			query = query + ` AND DATE(b.created_date) BETWEEN '` + startDate + `' AND '` + endDate + `'`
			queryT = queryT + ` AND DATE(b.created_date) BETWEEN '` + startDate + `' AND '` + endDate + `'`
		}

	}
	if isTransportation == true && isExperience == false {
		query = query + ` AND b.trans_id != '' `
		queryT = queryT + ` AND b.trans_id != '' `
	} else if isExperience == true && isTransportation == false {
		query = query + ` AND b.exp_id != '' `
		queryT = queryT + ` AND b.exp_id != '' `
	}
	unionQuery := query + ` UNION ` + queryT
	if limit != nil && offset != nil {
		unionQuery = unionQuery +
			` ORDER BY booking_date DESC LIMIT ` + strconv.Itoa(*limit) +
			` OFFSET ` + strconv.Itoa(*offset) + ` `

	}
	list, err := t.fetchWithJoin(ctx, unionQuery)
	if status != "" {
		if status == "pending" {
			transactionStatus = 0
		} else if status == "waitingApproval" {
			transactionStatus = 1
		} else if status == "confirm" {
			transactionStatus = 2
		}
		querySt := query + ` AND t.status = ` + strconv.Itoa(transactionStatus)
		queryTSt := queryT + ` AND t.status = ` + strconv.Itoa(transactionStatus)
		unionQuery = querySt + ` UNION ` + queryTSt
		if limit != nil && offset != nil {
			unionQuery = unionQuery +
				` ORDER BY booking_date DESC LIMIT ` + strconv.Itoa(*limit) +
				` OFFSET ` + strconv.Itoa(*offset) + ` `

		}
		list, err = t.fetchWithJoin(ctx, unionQuery)

		if status == "failed" {
			transactionStatus = 3
			cancelledStatus := 4
			querySt = query + ` AND t.status IN (` + strconv.Itoa(transactionStatus) + `,` + strconv.Itoa(cancelledStatus) + `)`
			queryTSt = queryT + ` AND t.status IN (` + strconv.Itoa(transactionStatus) + `,` + strconv.Itoa(cancelledStatus) + `)`
			unionQuery = querySt + ` UNION ` + queryTSt
			if limit != nil && offset != nil {
				unionQuery = unionQuery +
					` ORDER BY booking_date DESC LIMIT ` + strconv.Itoa(*limit) +
					` OFFSET ` + strconv.Itoa(*offset) + ` `

			}
			list, err = t.fetchWithJoin(ctx, unionQuery)
		}

		if status == "boarded" {
			transactionStatus = 2
			bookingStatus = 3
			querySt = query + ` AND t.status = ` + strconv.Itoa(transactionStatus) + ` AND b.status = ` + strconv.Itoa(bookingStatus)
			queryTSt = queryT + ` AND t.status = ` + strconv.Itoa(transactionStatus) + ` AND b.status = ` + strconv.Itoa(bookingStatus)
			unionQuery = querySt + ` UNION ` + queryTSt
			if limit != nil && offset != nil {
				unionQuery = unionQuery +
					` ORDER BY booking_date DESC LIMIT ` + strconv.Itoa(*limit) +
					` OFFSET ` + strconv.Itoa(*offset) + ` `

			}
			list, err = t.fetchWithJoin(ctx, unionQuery)
		}
	}
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t transactionRepository) fetchWithJoin(ctx context.Context, query string, args ...interface{}) ([]*models.TransactionOut, error) {
	rows, err := t.Conn.QueryContext(ctx, query, args...)
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

	result := make([]*models.TransactionOut, 0)
	for rows.Next() {
		t := new(models.TransactionOut)
		err = rows.Scan(
			&t.TransactionId,
			&t.ExpId,
			&t.ExpType,
			&t.ExpTitle,
			&t.BookingExpId,
			&t.BookingCode,
			&t.BookingDate,
			&t.CheckInDate,
			&t.BookedBy,
			&t.GuestDesc,
			&t.Email,
			&t.TransactionStatus,
			&t.BookingStatus,
			&t.TotalPrice,
			&t.ExperiencePaymentId,
			&t.MerchantName,
			&t.OrderId,
			&t.ExpDuration,
			&t.ProvinceName,
			&t.CountryName,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
func (t transactionRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.TransactionWMerchant, error) {
	rows, err := t.Conn.QueryContext(ctx, query, args...)
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

	result := make([]*models.TransactionWMerchant, 0)
	for rows.Next() {
		t := new(models.TransactionWMerchant)
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
			&t.BookingType,
			&t.BookingExpId,
			&t.PromoId,
			&t.PaymentMethodId,
			&t.ExperiencePaymentId,
			&t.Status,
			&t.TotalPrice,
			&t.Currency,
			&t.OrderId,
			&t.VaNumber,
			&t.ExChangeRates,
			&t.ExChangeCurrency,
			&t.MerchantId,
			&t.OrderIdBook,
			&t.BookedBy,
			&t.ExpTitle,
			&t.BookingDate,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (t transactionRepository) CountSuccess(ctx context.Context) (int, error) {
	query := `SELECT count(*) as count FROM transactions WHERE is_deleted = 0 AND is_active = 1 AND status = 2`

	rows, err := t.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	count, err := checkCount(rows)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return count, nil
}

func (t transactionRepository) Count(ctx context.Context, startDate, endDate, search, status string, merchantId string) (int, error) {
	query := `
	SELECT
		count(*) as count
	FROM 
		transactions t
		JOIN booking_exps b ON t.booking_exp_id = b.id
		JOIN experiences e ON b.exp_id = e.id
	WHERE
		t.is_deleted = 0
		AND t.is_active = 1`

	queryT := `
	SELECT
		count(*) as count
	FROM
		transactions t
		JOIN booking_exps b ON t.booking_exp_id = b.id
		JOIN transportations tr ON b.trans_id = tr.id
	WHERE
		t.is_deleted = 0
		AND t.is_active = 1`

	if merchantId != "" {
		query = query + ` AND e.merchant_id = '` + merchantId + `' `
		queryT = queryT + ` AND tr.merchant_id = '` + merchantId + `' `
	}
	if search != "" {
		keyword := `'%` + search + `%'`
		query = query + ` AND (LOWER(b.booked_by) LIKE LOWER(` + keyword + `) OR LOWER(b.order_id) LIKE LOWER(` + keyword + `))`
		queryT = queryT + ` AND (LOWER(b.booked_by) LIKE LOWER(` + keyword + `) OR LOWER(b.order_id) LIKE LOWER(` + keyword + `))`
	}
	if startDate != "" && endDate != "" {
		query = query + ` AND DATE(b.created_date) BETWEEN '` + startDate + `' AND '` + endDate + `'`
		queryT = queryT + ` AND DATE(b.created_date) BETWEEN '` + startDate + `' AND '` + endDate + `'`
	}
	unionQuery := query + ` UNION` + queryT
	rows, err := t.Conn.QueryContext(ctx, unionQuery)
	var transactionStatus int
	if status != "" {
		if status == "pending" {
			transactionStatus = 0
		} else if status == "waitingApproval" {
			transactionStatus = 1
		} else if status == "confirm" {
			transactionStatus = 2
		}
		querySt := query + ` AND t.status = ` + strconv.Itoa(transactionStatus)
		queryTSt := queryT + ` AND t.status = ` + strconv.Itoa(transactionStatus)
		unionQuery = querySt + ` UNION ` + queryTSt
		rows, err = t.Conn.QueryContext(ctx, unionQuery)

		if status == "failed" {
			transactionStatus = 3
			cancelledStatus := 4
			querySt = query + ` AND t.status IN (` + strconv.Itoa(transactionStatus) + `,` + strconv.Itoa(cancelledStatus) + `)`
			queryTSt = queryT + ` AND t.status IN (` + strconv.Itoa(transactionStatus) + `,` + strconv.Itoa(cancelledStatus) + `)`
			unionQuery = querySt + ` UNION ` + queryTSt
			rows, err = t.Conn.QueryContext(ctx, unionQuery)
		}

		if status == "boarded" {
			transactionStatus = 1
			bookingStatus := 3
			querySt = query + ` AND t.status = ` + strconv.Itoa(transactionStatus) + ` AND b.status = ` + strconv.Itoa(bookingStatus)
			queryTSt = queryT + ` AND t.status = ` + strconv.Itoa(transactionStatus) + ` AND b.status = ` + strconv.Itoa(bookingStatus)
			unionQuery = querySt + ` UNION ` + queryTSt
			rows, err = t.Conn.QueryContext(ctx, unionQuery)
		}
		if err != nil {
			logrus.Error(err)
			return 0, err
		}
	}

	count, err := checkCountUnion(rows)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return count, nil
}

func (m transactionRepository) GetById(ctx context.Context, id string) (*models.TransactionWMerchant, error) {
	query := `SELECT t.*,e.merchant_id,b.order_id as order_id_book,b.booked_by,e.exp_title,b.booking_date FROM transactions t
				join booking_exps b on t.booking_exp_id = b.id
				join experiences e on b.exp_id = e.id WHERE t.id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res := list[0]
		return res, nil
	} else {
		return nil, models.ErrNotFound
	}
	return nil, nil
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

func checkCountUnion(rows *sql.Rows) (result int, err error) {
	var count int
	results := make([]int, 2)
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
		results = append(results, count)
	}

	if len(results) > 0 {
		for _, r := range results {
			result += r
		}
	}
	return result, nil
}
