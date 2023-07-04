package controller

type signInPayload struct {
	Email    string `json:"email"    binding:"required" ja:"メールアドレス" example:"admin@menergia.co.jp"`
	Password string `json:"password" binding:"required" ja:"パスワード"   example:"Pass#123"`
}

type studentCreatePayload struct {
	FirstName    string `json:"first_name" binding:"required" example:"togo"`
	LastName     string `json:"last_name" binding:"required" example:"enomoto"`
	Grade        string `json:"grade" binding:"required" example:"1"`
	ManavisCode  string `json:"manavis_code" binding:"required" example:"200828"`
	IsHighSchool *bool  `json:"is_high_school" binding:"required" example:"true"`
}

type studentIDPathParams struct {
	ID string `uri:"id" binding:"required,uuid" example:"41265e3c-45ec-4264-9b32-3ffa1d08ebd3"`
}

type studentByManavisCodePayload struct {
	ManavisCode string `json:"manavis_code" binding:"required" example:"200828"`
}

type studentsByGradeAndIsInHighPayload struct {
	IsHighSchool *bool `json:"is_high_school" binding:"required" example:"true"`
	PageNumber   int   `json:"page_number" example:"1"`
}

type studentsByGradeAndIsInHighPathParams struct {
	Grade string `uri:"grade" binding:"required" example:"200828"`
}

type studentsWhoHadCheckedInWithDayOffestPayload struct {
	DayOffset int `json:"day_offset" binding:"required" example:"1"`
}
