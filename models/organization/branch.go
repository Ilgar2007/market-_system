package models

type Branch struct {
	Id          string `json:"id"`
	BranchCode  string `json:"branch_code"`
	Name        string `json:"name"`
	Address     string `json:"addres"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateBranch struct {
	BranchCode  string `json:"branch_code"`
	Name        string `json:"name"`
	Address     string `json:"addres"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateBranch struct {
	Id          string `json:"id"`
	BranchCode  string `json:"branch_code"`
	Name        string `json:"name"`
	Address     string `json:"addres"`
	PhoneNumber string `json:"phone_number"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type GetListBranchRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListBranchResponse struct {
	Count    int       `json:"count"`
	Branches []*Branch `json:"branches"`
}

// Филиал-Код
// 	MP
// 		SD
// Название
// 	Mega Planeta
// 		Samarqand Darvoza
// Аддрес
// 	Yunusobod
// Телефон
