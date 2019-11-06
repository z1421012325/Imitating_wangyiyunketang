package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func InstructorInfo(c *gin.Context){

	res := service.InstructorInfoService(c)

	c.JSON(200,res)
}
