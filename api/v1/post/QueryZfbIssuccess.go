package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)



func QueryZfbIssuccess(c *gin.Context)  {
	var service user.QueryZfbIssuccessService
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.QueryZfbIssuccess(c)
	c.JSON(200,res)

}
