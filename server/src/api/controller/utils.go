package controller

import (
	"github.com/gin-gonic/gin"
)

func MustBindQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		c.Abort()
		return c.Error(err).SetType(gin.ErrorTypeBind)
	}
	return nil
}
