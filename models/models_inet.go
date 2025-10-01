package models


type RegisterRequest struct {
	Username      string `json:"username" validate:"required,username"`
	Email         string `json:"email" validate:"required,email,min=3,max=32"`
	Password      string `json:"password" validate:"required,min=6,max=32"`
	LineID        string `json:"line_id" validate:"required,line_id"`
	Phone         string `json:"phone" validate:"required,numeric,min=9,max=10"`
	BussinessType string `json:"bussiness_type" validate:"required,bussiness_type"`
	Website       string `json:"website" validate:"required,subdomain"`
}
