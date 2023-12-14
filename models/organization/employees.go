package models

type Employee struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Branch      string `json:"branch"`
	SaleCenter  string `json:"sale_center"`
	UserType    string `json:"user_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateEmployee struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Branch      string `json:"branch"`
	SaleCenter  string `json:"sale_center"`
	UserType    string `json:"user_type"`
}
type UpdateEmployee struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Branch      string `json:"branch"`
	SaleCenter  string `json:"sale_center"`
	UserType    string `json:"user_type"`
}
type EmployeePrimaryKey struct {
	Id string `json:"id"`
}

type EmployeeGetListRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type EmployeeGetListResponse struct {
	Count     string      `json:"count"`
	Employees []*Employee `json:"sale_centers"`
}
