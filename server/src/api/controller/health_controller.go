package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Register(r gin.IRouter)
}

type healthController struct{}

func NewHealthController() HealthController {
	return &healthController{}
}

func (c *healthController) Register(r gin.IRouter) {
	r.GET("/health")
}

func (c *healthController) exec(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "health checked!",
		},
	)
}
