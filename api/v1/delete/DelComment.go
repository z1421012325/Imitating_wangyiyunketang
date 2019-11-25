package delete

import (
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

func DelComment(c *gin.Context){
	var service service.DelCommentService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.DelComment(c)
	c.JSON(200,res)
}