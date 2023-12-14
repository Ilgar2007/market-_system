package postgres

import (
	"context"
	"database/sql"
	"fmt"
	models "market/models/organization"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type centerRepo struct {
	db *pgxpool.Pool
}

func NewCenterRepo(db *pgxpool.Pool) *centerRepo {
	return &centerRepo{
		db: db,
	}
}

func (r *centerRepo) Create(ctx context.Context, req models.CreateSaleCenter) (*models.SaleCenter, error) {

	var (
		scId  = uuid.New().String()
		query = `
			INSERT INTO sale_center (
				"id",
				"name",
				"branch",
				"created_at"
			) VALUES ($1, $2, $3, NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		scId,
		req.Name,
		req.Branch,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, models.SaleCenterPrimaryKey{Id: scId})
}

func (r *centerRepo) GetByID(ctx context.Context, req models.SaleCenterPrimaryKey) (*models.SaleCenter, error) {

	var whereField = "id"

	// if len(req.Login) > 0 {
	// 	whereField = "login"
	// }
	query := `
			SELECT
				"id",
				"name",
				"branch",
				"created_at",
				"updated_at"
			FROM "sale_center"
			WHERE ` + whereField + ` = $1
		`

	var (
		id         sql.NullString
		name       sql.NullString
		branch     sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&branch,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.SaleCenter{
		Id:        id.String,
		Name:      name.String,
		Branch:    branch.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (r *centerRepo) Delete(ctx context.Context, req models.SaleCenterPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM sale_center WHERE id = $1", req.Id)
	return err
}

func (r *centerRepo) Update(ctx context.Context, req models.UpdateSaleCenter) (int64, error) {

	query := `
		UPDATE "sale_center"
			SET
			"name" = $2,
			"branch" = $3
		WHERE id = $1
	`
	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.Branch,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *centerRepo) GetList(ctx context.Context, req models.SaleCenterGetListRequest) (*models.SaleCenterGetListResponse, error) {
	var (
		resp   models.SaleCenterGetListResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"

		Id        sql.NullString
		Name      sql.NullString
		Branch    sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
		query     = `
			SELECT
				COUNT(*) OVER(),
				"id",
				"name",
				"branch",
				"created_at",
				"updated_at"
			FROM "sale_center"
		`
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	fmt.Println(query)
	for rows.Next() {
		var sale_center models.SaleCenter

		err := rows.Scan(
			&resp.Count,
			&Id,
			&Name,
			&Branch,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sale_center = models.SaleCenter{
			Id:        Id.String,
			Branch:    Branch.String,
			Name:      Name.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		}
		resp.SaleCenters = append(resp.SaleCenters, &sale_center)
	}
	return &resp, nil
}
