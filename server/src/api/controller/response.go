package controller

type studentResponse struct {
	ID           string `json:"id" example:"41265e3c-45ec-4264-9b32-3ffa1d08ebd3"`
	FirstName    string `json:"first_name" example:"togo"`
	LastName     string `json:"last_name" example:"enomoto"`
	Grade        int    `json:"grade" example:"1"`
	ManavisCode  string `json:"manavis_code" example:"200828"`
	IsHighSchool bool   `json:"is_high_school" example:"true"`
}

type paginationResponse struct {
	Page  int `json:"page"  binding:"required" example:"1"`
	Size  int `json:"size"  binding:"required" example:"50"`
	Total int `json:"total" binding:"required" example:"150"`
}

type studentsResponse struct {
	Students []studentResponse
	paginationResponse
}
