package auth

import "github.com/gin-gonic/gin"

type Controller interface {
	Login(c *gin.Context)
	Validate(c *gin.Context)
}
