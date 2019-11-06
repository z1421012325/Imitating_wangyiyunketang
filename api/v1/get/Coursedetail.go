package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Coursedetail(c *gin.Context){
	res := service.CoursedetailService(c)
	c.JSON(200,res)
}