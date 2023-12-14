package models

type Provider struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Active      bool   `json:"active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateProvider struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Active      bool   `json:"active"`
}

type UpdateProvider struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Active      bool   `json:"active"`
}

type ProviderPrimaryKey struct {
	Id string `json:"id"`
}

type ProviderGetListRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type ProviderGetListResponse struct {
	Count     string      `json:"count"`
	Providers []*Provider `json:"sale_centers"`
}
