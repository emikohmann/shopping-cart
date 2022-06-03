package items

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Search(c *gin.Context)
}
