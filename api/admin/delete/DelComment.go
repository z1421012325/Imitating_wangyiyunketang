package delete

import (
	"demos/serialize"
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func DelComment(c *gin.Context){
	var service admin.DelCommentServic
	if err := c.ShouldBind(&service); err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.DelComment(c)
	c.JSON(200,res)
}
