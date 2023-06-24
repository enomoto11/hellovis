package controller

import (
	"fmt"
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
	var req studentCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade, gradeErr := strconv.Atoi(req.Grade)
	if gradeErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": gradeErr.Error()})
		return
	}
	params := &service.CreateStudentParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Grade:        grade,
		IsHighSchool: *req.IsHighSchool,
		ManavisCode:  req.ManavisCode,
	}
	_, err := c.studentService.Create(ctx, params)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{"error": err.Error()})
		return
	} else {
		fmt.Println("success")
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "new member has been created!",
	})
}
