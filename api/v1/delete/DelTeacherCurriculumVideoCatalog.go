package delete

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func DelTeacherCurriculumVideoCatalog(c *gin.Context){
	var service user.DelTeacherCurriculumVideoCatalogService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.DelTeacherCurriculumVideoCatalog(c)
	c.JSON(200,res)
}
