package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowComment(c *gin.Context)  {
	res := admin.ShowCommentService(c)
	c.JSON(200,res)
}
