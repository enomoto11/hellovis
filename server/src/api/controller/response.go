package controller

type StudentResponse struct {
	ID           string `json:"id" example:"41265e3c-45ec-4264-9b32-3ffa1d08ebd3"`
	FirstName    string `json:"first_name" example:"togo"`
	LastName     string `json:"last_name" example:"enomoto"`
	Grade        int    `json:"grade" example:"1"`
	ManavisCode  string `json:"manavis_code" example:"200828"`
	IsHighSchool bool   `json:"is_high_school" example:"true"`
}
