package postgres

import (
	"context"
	"fmt"
	models "market/models/organization"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type paymentRepo struct {
	db *pgxpool.Pool
}

func NewPaymentRepo(db *pgxpool.Pool) *paymentRepo {
	return &paymentRepo{
		db: db,
	}
}

func (r *paymentRepo) Create(ctx context.Context, req *models.CreatePayment) (*models.Payment, error) {
	paymentId := uuid.New().String()
	query := `
		INSERT INTO payment (
			id,
			sale_id,
			cash,
			uzcard,
			payme,
			click,
			humo,
			apelsin,
			total_amount,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`

	_, err := r.db.Exec(ctx,
		query,
		paymentId,
		req.SaleID,
		req.Cash,
		req.Uzcard,
		req.Payme,
		req.Click,
		req.Humo,
		req.Apelsin,
		req.TotalAmount,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, &models.PaymentPrimaryKey{Id: paymentId})
}

func (r *paymentRepo) GetByID(ctx context.Context, req *models.PaymentPrimaryKey) (*models.Payment, error) {
	query := `
		SELECT
			id,
			sale_id,
			cash,
			uzcard,
			payme,
			click,
			humo,
			apelsin,
			total_amount,
			created_at,
			updated_at
		FROM payment
		WHERE id = $1
	`

	var payment models.Payment
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&payment.Id,
		&payment.SaleID,
		&payment.Cash,
		&payment.Uzcard,
		&payment.Payme,
		&payment.Click,
		&payment.Humo,
		&payment.Apelsin,
		&payment.TotalAmount,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *paymentRepo) GetList(ctx context.Context, req *models.GetListPaymentRequest) (*models.GetListPaymentResponse, error) {
	var (
		resp   models.GetListPaymentResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += fmt.Sprintf(" AND sale_id ILIKE '%%%s%%'", req.Search)
	}

	if len(req.Query) > 0 {
		where += req.Query
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			sale_id,
			cash,
			uzcard,
			payme,
			click,
			humo,
			apelsin,
			total_amount,
			created_at,
			updated_at
		FROM payment
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var payment models.Payment
		err = rows.Scan(
			&resp.Count,
			&payment.Id,
			&payment.SaleID,
			&payment.Cash,
			&payment.Uzcard,
			&payment.Payme,
			&payment.Click,
			&payment.Humo,
			&payment.Apelsin,
			&payment.TotalAmount,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Payments = append(resp.Payments, &payment)
	}

	return &resp, nil
}

func (r *paymentRepo) Update(ctx context.Context, req *models.UpdatePayment) (int64, error) {
	query := `
		UPDATE payment
		SET
			cash = $2,
			uzcard = $3,
			payme = $4,
			click = $5,
			humo = $6,
			apelsin = $7,
			total_amount = $8,
			updated_at = NOW()
		WHERE id = $1
	`

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Cash,
		req.Uzcard,
		req.Payme,
		req.Click,
		req.Humo,
		req.Apelsin,
		req.TotalAmount,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *paymentRepo) Delete(ctx context.Context, req *models.PaymentPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM payment WHERE id = $1", req.Id)
	return err
}
