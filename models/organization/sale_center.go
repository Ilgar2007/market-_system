package models

type SaleCenter struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Branch    string `json:"branch"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateSaleCenter struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
}

type UpdateSaleCenter struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Branch string `json:"branch"`
}

type SaleCenterPrimaryKey struct {
	Id string `json:"id"`
}

type SaleCenterGetListRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type SaleCenterGetListResponse struct {
	Count       string        `json:"count"`
	SaleCenters []*SaleCenter `json:"sale_centers"`
}
