package storage

import (
	"context"
	models "market/models/organization"
)

type StorageI interface {
	Branch() BranchRepoI
	Center() CenterRepoI
	Provider() ProviderRepoI
	Employee() EmployeeRepoI
	Category() CategoryRepoI
	User() UserRepoI
	Product() ProductRepoI
	Income() IncomeRepoI
	IncomeProduct() IncomeProductRepoI
	Remainder() RemainderRepoI
	Sale() SaleRepoI
	Sale_Product() SaleProductRepoI
	Payment() PaymentRepoI
	Transaction() TransactionRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req models.CreateBranch) (*models.Branch, error)
	GetByID(ctx context.Context, req models.BranchPrimaryKey) (*models.Branch, error)
	GetList(ctx context.Context, req models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(ctx context.Context, req models.UpdateBranch) (int64, error)
	Delete(ctx context.Context, req models.BranchPrimaryKey) error
}
type CenterRepoI interface {
	Create(ctx context.Context, req models.CreateSaleCenter) (*models.SaleCenter, error)
	GetByID(ctx context.Context, req models.SaleCenterPrimaryKey) (*models.SaleCenter, error)
	GetList(ctx context.Context, req models.SaleCenterGetListRequest) (*models.SaleCenterGetListResponse, error)
	Update(ctx context.Context, req models.UpdateSaleCenter) (int64, error)
	Delete(ctx context.Context, req models.SaleCenterPrimaryKey) error
}
type ProviderRepoI interface {
	Create(ctx context.Context, req models.CreateProvider) (*models.Provider, error)
	GetByID(ctx context.Context, req models.ProviderPrimaryKey) (*models.Provider, error)
	GetList(ctx context.Context, req models.ProviderGetListRequest) (*models.ProviderGetListResponse, error)
	Update(ctx context.Context, req models.UpdateProvider) (int64, error)
	Delete(ctx context.Context, req models.ProviderPrimaryKey) error
}

type EmployeeRepoI interface {
	Create(ctx context.Context, req models.CreateEmployee) (*models.Employee, error)
	GetByID(ctx context.Context, req models.EmployeePrimaryKey) (*models.Employee, error)
	GetList(ctx context.Context, req models.EmployeeGetListRequest) (*models.EmployeeGetListResponse, error)
	Update(ctx context.Context, req models.UpdateEmployee) (int64, error)
	Delete(ctx context.Context, req models.EmployeePrimaryKey) error
}

type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (*models.Product, error)
	GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) error
}

type IncomeRepoI interface {
	Create(ctx context.Context, req *models.CreateIncome) (*models.Income, error)
	GetByID(ctx context.Context, req *models.IncomePrimaryKey) (*models.Income, error)
	GetList(ctx context.Context, req *models.GetListIncomeRequest) (*models.GetListIncomeResponse, error)
	Update(ctx context.Context, req *models.UpdateIncome) (int64, error)
	Delete(ctx context.Context, req *models.IncomePrimaryKey) error
}

type IncomeProductRepoI interface {
	Create(ctx context.Context, req *models.CreateIncomeProduct) (*models.IncomeProduct, error)
	GetByID(ctx context.Context, req *models.IncomeProductPrimaryKey) (*models.IncomeProduct, error)
	GetList(ctx context.Context, req *models.GetListIncomeProductRequest) (*models.GetListIncomeProductResponse, error)
	Update(ctx context.Context, req *models.UpdateIncomeProduct) (int64, error)
	Delete(ctx context.Context, req *models.IncomeProductPrimaryKey) error
}

type RemainderRepoI interface {
	Create(ctx context.Context, req *models.CreateRemainder) (*models.Remainder, error)
	GetByID(ctx context.Context, req *models.RemainderPrimaryKey) (*models.Remainder, error)
	GetList(ctx context.Context, req *models.GetListRemainderRequest) (*models.GetListRemainderResponse, error)
	Update(ctx context.Context, req *models.UpdateRemainder) (int64, error)
	Delete(ctx context.Context, req *models.RemainderPrimaryKey) error
}

type SaleRepoI interface {
	Create(ctx context.Context, req *models.CreateSale) (*models.Sale, error)
	GetByID(ctx context.Context, req *models.SalePrimaryKey) (*models.Sale, error)
	GetList(ctx context.Context, req *models.GetListSaleRequest) (*models.GetListSaleResponse, error)
	Update(ctx context.Context, req *models.UpdateSale) (int64, error)
	Delete(ctx context.Context, req *models.SalePrimaryKey) error
}

type SaleProductRepoI interface {
	Create(ctx context.Context, req *models.CreateSaleProduct) (*models.SaleProduct, error)
	GetByID(ctx context.Context, req *models.SaleProductPrimaryKey) (*models.SaleProduct, error)
	GetList(ctx context.Context, req *models.GetListSaleProductRequest) (*models.GetListSaleProductResponse, error)
	Update(ctx context.Context, req *models.UpdateSaleProduct) (int64, error)
	Delete(ctx context.Context, req *models.SaleProductPrimaryKey) error
}

type PaymentRepoI interface {
	Create(ctx context.Context, req *models.CreatePayment) (*models.Payment, error)
	GetByID(ctx context.Context, req *models.PaymentPrimaryKey) (*models.Payment, error)
	GetList(ctx context.Context, req *models.GetListPaymentRequest) (*models.GetListPaymentResponse, error)
	Update(ctx context.Context, req *models.UpdatePayment) (int64, error)
	Delete(ctx context.Context, req *models.PaymentPrimaryKey) error
}

type TransactionRepoI interface {
	Create(ctx context.Context, req *models.CreateTransaction) (*models.Transaction, error)
	GetByID(ctx context.Context, req *models.TransactionPrimaryKey) (*models.Transaction, error)
	GetList(ctx context.Context, req *models.GetListTransactonRequest) (*models.GetListTransactionResponse, error)
	Update(ctx context.Context, req *models.UpdateTransaction) (int64, error)
	Delete(ctx context.Context, req *models.TransactionPrimaryKey) error
}
type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.User, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) error
}
type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (*models.Category, error)
	GetByID(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) error
}
