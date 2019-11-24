package post

import (
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

func AddCollection(c *gin.Context){
	var service service.AddCollectionService
	if err := c.ShouldBind(&service); err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.AddCollection(c)
	c.JSON(200,res)
}
