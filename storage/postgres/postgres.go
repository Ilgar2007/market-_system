package postgres

import (
	"context"
	"fmt"
	"market/config"
	"market/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db             *pgxpool.Pool
	branch         storage.BranchRepoI
	category       storage.CategoryRepoI
	center         storage.CenterRepoI
	provider       storage.ProviderRepoI
	employee       storage.EmployeeRepoI
	product        storage.ProductRepoI
	income         storage.IncomeRepoI
	income_product storage.IncomeProductRepoI
	remainder      storage.RemainderRepoI
	sale           storage.SaleRepoI
	sale_product   storage.SaleProductRepoI
	payment        storage.PaymentRepoI
	user           storage.UserRepoI
	transaction    storage.TransactionRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
			cfg.PostgresHost,
			cfg.PostgresUser,
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
		),
	)

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pgxpool,
	}, nil
}
func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}
func (s *Store) Branch() storage.BranchRepoI {

	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}
func (s *Store) Center() storage.CenterRepoI {

	if s.center == nil {
		s.center = NewCenterRepo(s.db)
	}

	return s.center
}

func (s *Store) Provider() storage.ProviderRepoI {
	if s.provider == nil {
		s.provider = NewProviderRepo(s.db)
	}

	return s.provider
}
func (s *Store) Employee() storage.EmployeeRepoI {
	if s.employee == nil {
		s.employee = NewEmployeeRepo(s.db)
	}

	return s.employee
}

func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) Income() storage.IncomeRepoI {

	if s.income == nil {
		s.income = NewIncomeRepo(s.db)
	}

	return s.income
}

func (s *Store) IncomeProduct() storage.IncomeProductRepoI {

	if s.income_product == nil {
		s.income_product = NewIncomeProductRepo(s.db)
	}

	return s.income_product
}

func (s *Store) Remainder() storage.RemainderRepoI {

	if s.remainder == nil {
		s.remainder = NewRemainderRepo(s.db)
	}

	return s.remainder
}

func (s *Store) Sale() storage.SaleRepoI {

	if s.sale == nil {
		s.sale = NewSaleRepo(s.db)
	}

	return s.sale
}

func (s *Store) Sale_Product() storage.SaleProductRepoI {

	if s.sale_product == nil {
		s.sale_product = NewSaleProductRepo(s.db)
	}

	return s.sale_product
}

func (s *Store) Payment() storage.PaymentRepoI {

	if s.payment == nil {
		s.payment = NewPaymentRepo(s.db)
	}

	return s.payment
}

func (s *Store) Transaction() storage.TransactionRepoI {

	if s.transaction == nil {
		s.transaction = NewTransactionRepo(s.db)
	}

	return s.transaction
}
func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
