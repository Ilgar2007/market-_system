package postgres

import (
	"context"
	"database/sql"
	"fmt"
	models "market/models/organization"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(ctx context.Context, req models.CreateBranch) (*models.Branch, error) {

	var (
		brId  = uuid.New().String()
		query = `
			INSERT INTO branch (
				"id",
				"branch_code",
				"name",
				"address",
				"phone_number",
				"created_at"
			) VALUES ($1, $2, $3, $4, $5,NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		brId,
		req.BranchCode,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, models.BranchPrimaryKey{Id: brId})
}

func (r *branchRepo) GetByID(ctx context.Context, req models.BranchPrimaryKey) (*models.Branch, error) {

	var whereField = "id"

	// if len(req.Login) > 0 {
	// 	whereField = "login"
	// }
	query := `
			SELECT
				"id",
				"branch_code",
				"name",
				"address",
				"phone_number",
				"created_at",
				"updated_at"
			FROM "branch"
			WHERE ` + whereField + ` = $1
		`

	var (
		id           sql.NullString
		branch_code  sql.NullString
		name         sql.NullString
		address      sql.NullString
		phone_number sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&branch_code,
		&name,
		&address,
		&phone_number,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:          id.String,
		BranchCode:  branch_code.String,
		Name:        name.String,
		Address:     address.String,
		PhoneNumber: phone_number.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}, nil
}

func (r *branchRepo) Delete(ctx context.Context, req models.BranchPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM branch WHERE id = $1", req.Id)
	return err
}

func (r *branchRepo) Update(ctx context.Context, req models.UpdateBranch) (int64, error) {

	query := `
		UPDATE "branch"
			SET
			"branch_code" = $2,
			"name" = $3,
			"address" = $4,
			"phone_number" = $5
		WHERE id = $1
	`
	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.BranchCode,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *branchRepo) GetList(ctx context.Context, req models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	var (
		resp   models.GetListBranchResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"

		Id          sql.NullString
		BranchCode  sql.NullString
		Name        sql.NullString
		Address     sql.NullString
		PhoneNumber sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
		query       = `
			SELECT
				COUNT(*) OVER(),
				"id",
				"branch_code",
				"name",
				"address",
				"phone_number",
				"created_at",
				"updated_at"
			FROM "branch"
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
		var branch models.Branch

		err := rows.Scan(
			&resp.Count,
			&Id,
			&BranchCode,
			&Name,
			&Address,
			&PhoneNumber,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		branch = models.Branch{
			Id:          Id.String,
			BranchCode:  BranchCode.String,
			Name:        Name.String,
			Address:     Address.String,
			PhoneNumber: PhoneNumber.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		}
		resp.Branches = append(resp.Branches, &branch)
	}
	return &resp, nil
}
