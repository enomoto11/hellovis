package controller

type studentCreateRequest struct {
	FirstName    string `json:"first_name" binding:"required" example:"togo"`
	LastName     string `json:"last_name" binding:"required" example:"enomoto"`
	Grade        string `json:"grade" binding:"required" example:"1"`
	ManavisCode  string `json:"manavis_code" binding:"required" example:"200828"`
	IsHighSchool *bool  `json:"is_high_school" binding:"required" example:"true"`
}
