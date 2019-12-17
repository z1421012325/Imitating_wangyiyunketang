package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func RecoveryCurriculum(c *gin.Context){

	var service user.RecoveryCurriculumService
	if err := c.ShouldBind(&service); err != nil{
		c.JSON(200,serialize.ParamErr("",nil))
		return
	}

	res := service.RecoveryCurriculum(c)
	c.SecureJSON(200,res)

}
