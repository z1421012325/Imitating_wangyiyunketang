package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ModifyTeacherCurriculumVideoCatalog(c *gin.Context){
	var service user.ModifyTeacherCurriculumVideoService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.ModifyTeacherCurriculumVideo(c)
	c.JSON(200,res)
}
