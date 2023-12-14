package postgres

import (
	"context"
	"database/sql"
	"fmt"
	models "market/models/organization"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type providerRepo struct {
	db *pgxpool.Pool
}

func NewProviderRepo(db *pgxpool.Pool) *providerRepo {
	return &providerRepo{
		db: db,
	}
}

func (r *providerRepo) Create(ctx context.Context, req models.CreateProvider) (*models.Provider, error) {

	var (
		prId  = uuid.New().String()
		query = `
			INSERT INTO provider (
				"id",
				"name",
				"phone_number",
				"active",
				"created_at"
			) VALUES ($1, $2, $3, $4,NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		prId,
		req.Name,
		req.PhoneNumber,
		req.Active,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, models.ProviderPrimaryKey{Id: prId})
}

func (r *providerRepo) GetByID(ctx context.Context, req models.ProviderPrimaryKey) (*models.Provider, error) {

	var whereField = "id"

	// if len(req.Login) > 0 {
	// 	whereField = "login"
	// }
	query := `
			SELECT
				"id",
				"name",
				"phone_number",
				"active",
				"created_at",
				"updated_at"
			FROM "provider"
			WHERE ` + whereField + ` = $1
		`

	var (
		id           sql.NullString
		name         sql.NullString
		phone_number sql.NullString
		active       sql.NullBool
		created_at   sql.NullString
		updated_at   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&phone_number,
		&active,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Provider{
		Id:          id.String,
		Name:        name.String,
		PhoneNumber: phone_number.String,
		Active:      active.Bool,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}, nil
}

func (r *providerRepo) Delete(ctx context.Context, req models.ProviderPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM provider WHERE id = $1", req.Id)
	return err
}

func (r *providerRepo) Update(ctx context.Context, req models.UpdateProvider) (int64, error) {

	query := `
		UPDATE "provider"
			SET
			"name" = $2,
			"phone_number" = $3,
			"active" = $4
		WHERE id = $1
	`
	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.PhoneNumber,
		req.Active,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}
func (r *providerRepo) GetList(ctx context.Context, req models.ProviderGetListRequest) (*models.ProviderGetListResponse, error) {
	var (
		resp   models.ProviderGetListResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"

		Id          sql.NullString
		Name        sql.NullString
		PhoneNumber sql.NullString
		Active      sql.NullBool
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
		query       = `
			SELECT
				COUNT(*) OVER(),
				"id",
				"name",
				"phone_number",
				"active",
				"created_at",
				"updated_at"
			FROM "provider"
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
		var provider models.Provider

		err := rows.Scan(
			&resp.Count,
			&Id,
			&Name,
			&PhoneNumber,
			&Active,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		provider = models.Provider{
			Id:          Id.String,
			Name:        Name.String,
			PhoneNumber: PhoneNumber.String,
			Active:      Active.Bool,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		}
		resp.Providers = append(resp.Providers, &provider)
	}
	return &resp, nil
}
