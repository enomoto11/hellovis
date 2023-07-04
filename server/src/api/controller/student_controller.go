package controller

import (
	"hellovis/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type StudentController interface {
	Register(r gin.IRouter)
}

type studentController struct {
	studentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &studentController{studentService}
}

func (c *studentController) Register(r gin.IRouter) {
	r.POST("/students", c.createStudent)
	r.GET("/students/:id", c.getStudentByID)
	r.GET("/students/manavis-code", c.getStudentByManaVisCode)
	r.GET("/students/grade/:grade", c.getAllByGradeAndIsInHigh)
	r.GET("students/checked-in", c.getAllWhoHadCheckedInWithDayOffest)
}

func (c *studentController) createStudent(ctx *gin.Context) {
	var payload studentCreatePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var validate = validator.New()
	if err := validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade, gradeErr := strconv.Atoi(payload.Grade)
	if gradeErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": gradeErr.Error()})
		return
	}
	params := &service.CreateStudentParams{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Grade:        grade,
		IsHighSchool: *payload.IsHighSchool,
		ManavisCode:  payload.ManavisCode,
	}
	_, err := c.studentService.Create(ctx, params)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "new member has been created!",
	})
}

func (c *studentController) getStudentByID(ctx *gin.Context) {
	var pathParams studentIDPathParams
	if pathParamsErr := ctx.ShouldBindUri(&pathParams); pathParamsErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": pathParamsErr.Error()})
		return
	}

	id, idErr := uuid.Parse(pathParams.ID)
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": idErr.Error()})
		return
	}

	result, err := c.studentService.FindByID(ctx, id)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{"error": err.Error()})
		return
	}
	res := studentResponse{
		ID:           result.GetID().String(),
		FirstName:    result.GetFirstName(),
		LastName:     result.GetLastName(),
		Grade:        result.GetGrade(),
		ManavisCode:  result.GetManavisCode(),
		IsHighSchool: result.GetIsInHighSchool(),
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *studentController) getStudentByManaVisCode(ctx *gin.Context) {
	var payload studentByManavisCodePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var validate = validator.New()
	if err := validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.studentService.FindByManavisCode(ctx, payload.ManavisCode)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{"error": err.Error()})
		return
	}
	res := studentResponse{
		ID:           result.GetID().String(),
		FirstName:    result.GetFirstName(),
		LastName:     result.GetLastName(),
		Grade:        result.GetGrade(),
		ManavisCode:  result.GetManavisCode(),
		IsHighSchool: result.GetIsInHighSchool(),
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *studentController) getAllByGradeAndIsInHigh(ctx *gin.Context) {
	var pathParams studentsByGradeAndIsInHighPathParams
	if pathParamsErr := ctx.ShouldBindUri(&pathParams); pathParamsErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": pathParamsErr.Error()})
		return
	}

	var payload studentsByGradeAndIsInHighPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var validate = validator.New()
	if err := validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade, gradeErr := strconv.Atoi(pathParams.Grade)
	if gradeErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": gradeErr.Error()})
		return
	}

	params := &service.FindAllByGradeAndIsInHighParams{
		Grade:        grade,
		IsHighSchool: *payload.IsHighSchool,
		PageNumber:   payload.PageNumber,
	}

	results, err := c.studentService.FindAllByGradeAndIsInHigh(ctx, params)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{"error": err.Error()})
		return
	}

	var students []studentResponse

	for _, result := range results.GetContents() {
		student := studentResponse{
			ID:           result.GetID().String(),
			FirstName:    result.GetFirstName(),
			LastName:     result.GetLastName(),
			Grade:        result.GetGrade(),
			ManavisCode:  result.GetManavisCode(),
			IsHighSchool: result.GetIsInHighSchool(),
		}
		students = append(students, student)
	}

	res := &studentsResponse{
		Students: students,
		paginationResponse: paginationResponse{
			Page:  results.GetPageNumber(),
			Size:  results.GetPageSize(),
			Total: results.GetTotalCount(),
		},
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *studentController) getAllWhoHadCheckedInWithDayOffest(ctx *gin.Context) {
	var payload studentsWhoHadCheckedInWithDayOffestPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var validate = validator.New()
	if err := validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := service.FindAllWhoHadCheckedInWithDayOffestParams{
		DayOffset: payload.DayOffset,
	}

	results, err := c.studentService.FindAllWhoHadCheckedInWithDayOffest(ctx, params)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{"error": err.Error()})
		return
	}

	var res []studentResponse
	for _, result := range results {
		student := studentResponse{
			ID:           result.GetID().String(),
			FirstName:    result.GetFirstName(),
			LastName:     result.GetLastName(),
			Grade:        result.GetGrade(),
			ManavisCode:  result.GetManavisCode(),
			IsHighSchool: result.GetIsInHighSchool(),
		}
		res = append(res, student)
	}

	ctx.JSON(http.StatusOK, res)
}
