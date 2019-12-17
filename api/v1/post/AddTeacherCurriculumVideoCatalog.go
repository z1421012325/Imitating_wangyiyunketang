package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func AddTeacherCurriculumVideoCatalog(c *gin.Context){
	var service user.AddTeacherCurriculumVideoCatalogService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.AddTeacherCurriculumVideoCatalog(c)
	c.JSON(200,res)
}
