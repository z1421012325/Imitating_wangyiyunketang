package delete

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func DelCollection(c *gin.Context){
	var service user.DelCollectionService
	if err := c.ShouldBind(&service); err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.DelCollection(c)
	c.JSON(200,res)

}
