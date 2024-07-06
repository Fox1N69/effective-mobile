package models

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserFilters struct {
	PassportNumber string `json:"passportNumber,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
