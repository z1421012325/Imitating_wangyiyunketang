package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func CourseDetail(c *gin.Context){
	res := service.CourseDetailService(c)
	c.JSON(200,res)
}