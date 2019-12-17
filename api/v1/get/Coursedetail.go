package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func CourseDetail(c *gin.Context){
	res := user.CourseDetailService(c)
	c.JSON(200,res)
}