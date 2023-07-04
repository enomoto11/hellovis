package controller

import (
	"hellovis/api/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(r gin.IRouter)
}

type authController struct {
	studentService service.StudentService
}

func NewAuthController(studentService service.StudentService) AuthController {
	return &authController{studentService}
}

func (c *authController) Register(r gin.IRouter) {
	r.POST("/signin", c.signin)
}

func (c *authController) signin(ctx *gin.Context) {}
