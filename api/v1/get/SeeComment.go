package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func SeeComment(c *gin.Context){
	res := service.SeeCommentService(c)
	c.SecureJSON(200,res)
}
