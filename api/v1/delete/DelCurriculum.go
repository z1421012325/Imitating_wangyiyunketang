package delete

import (
	"demos/serialize"
	"demos/service"
	"fmt"
	"github.com/gin-gonic/gin"
)



func DelCurriculum(c *gin.Context){
	var service service.DelCurriculumService
	fmt.Println("tag 检测之前")
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		fmt.Println("tag 检测出错")
		return
	}
	fmt.Println("tag 检测完成,进入下一步")
	res := service.DelCurriculum(c)
	c.JSON(200,res)
}
