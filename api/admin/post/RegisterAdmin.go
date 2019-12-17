package post

import (
	"demos/serialize"
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdmin(c *gin.Context){
	var service admin.RegisterAdminService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.RegisterAdmin()
	c.JSON(200,res)
}
