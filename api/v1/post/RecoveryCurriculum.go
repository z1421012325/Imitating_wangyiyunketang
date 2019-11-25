package post

import (
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

func RecoveryCurriculum(c *gin.Context){

	var service service2.RecoveryCurriculumService
	if err := c.ShouldBind(&service); err != nil{
		c.JSON(200,serialize.ParamErr("",nil))
		return
	}

	res := service.RecoveryCurriculum(c)
	c.SecureJSON(200,res)

}
