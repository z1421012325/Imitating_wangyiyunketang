package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func SeeComment(c *gin.Context){
	res := user.SeeCommentService(c)
	c.SecureJSON(200,res)
}
