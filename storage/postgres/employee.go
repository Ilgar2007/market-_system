package postgres

import (
	"context"
	"database/sql"
	"fmt"
	models "market/models/organization"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type employeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) Create(ctx context.Context, req models.CreateEmployee) (*models.Employee, error) {

	var (
		prId  = uuid.New().String()
		query = `
			INSERT INTO employee (
				"id",
				"name",
				"last_name",
				"phone_number",
				"login",
				"password",
				"branch",
				"sale_center",
				"user_type",
				"created_at"
			) VALUES ($1, $2, $3, $4, $5, $6 , $7 , $8 , $9 ,NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		prId,
		req.Name,
		req.LastName,
		req.PhoneNumber,
		req.Login,
		req.Password,
		req.Branch,
		req.SaleCenter,
		req.UserType,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, models.EmployeePrimaryKey{Id: prId})
}

func (r *employeeRepo) GetByID(ctx context.Context, req models.EmployeePrimaryKey) (*models.Employee, error) {

	var whereField = "id"

	// if len(req.Login) > 0 {
	// 	whereField = "login"
	// }
	query := `
			SELECT
				"id",
				"name",
				"last_name",
				"phone_number",
				"login",
				"password",
				"branch",
				"sale_center",
				"user_type",
				"created_at"
				"updated_at"
			FROM "employee"
			WHERE ` + whereField + ` = $1
		`

	var (
		id           sql.NullString
		name         sql.NullString
		last_name    sql.NullString
		phone_number sql.NullString
		login        sql.NullString
		password     sql.NullString
		branch       sql.NullString
		sale_center  sql.NullString
		user_type    sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&last_name,
		&phone_number,
		&login,
		&password,
		&branch,
		&sale_center,
		&user_type,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Employee{
		Id:          id.String,
		Name:        name.String,
		LastName:    last_name.String,
		PhoneNumber: phone_number.String,
		Login:       login.String,
		Password:    password.String,
		Branch:      branch.String,
		SaleCenter:  sale_center.String,
		UserType:    user_type.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}, nil
}

func (r *employeeRepo) Delete(ctx context.Context, req models.EmployeePrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM employee WHERE id = $1", req.Id)
	return err
}

func (r *employeeRepo) Update(ctx context.Context, req models.UpdateEmployee) (int64, error) {

	query := `
		UPDATE "employee"
			SET
				"id" = $2,
				"name" = $3,
				"last_name" = $4,
				"phone_number" = $5,
				"login" = $6,
				"password" = $7,
				"branch" = $8,
				"sale_center" = $9,
				"user_type" = $10,
		WHERE id = $1
	`
	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.LastName,
		req.PhoneNumber,
		req.Login,
		req.Password,
		req.Branch,
		req.SaleCenter,
		req.UserType,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}
func (r *employeeRepo) GetList(ctx context.Context, req models.EmployeeGetListRequest) (*models.EmployeeGetListResponse, error) {
	var (
		resp   models.EmployeeGetListResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"

		Id          sql.NullString
		Name        sql.NullString
		LastName    sql.NullString
		PhoneNumber sql.NullString
		Login       sql.NullString
		Password    sql.NullString
		Branch      sql.NullString
		SaleCenter  sql.NullString
		UserType    sql.NullString
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
		var employee models.Employee

		err := rows.Scan(
			&resp.Count,
			&Id,
			&Name,
			&LastName,
			&PhoneNumber,
			&Login,
			&Password,
			&Branch,
			&SaleCenter,
			&UserType,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employee = models.Employee{
			Id:          Id.String,
			Name:        Name.String,
			LastName:    LastName.String,
			PhoneNumber: PhoneNumber.String,
			Login:       Login.String,
			Password:    Password.String,
			Branch:      Branch.String,
			SaleCenter:  SaleCenter.String,
			UserType:    UserType.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		}
		resp.Employees = append(resp.Employees, &employee)
	}
	return &resp, nil
}
