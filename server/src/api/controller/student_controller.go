package controller

import (
	"hellovis/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
}

func (c *studentController) createStudent(ctx *gin.Context) {
	var payload studentCreateRequest
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
		IsHighSchool: payload.IsHighSchool,
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
