package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func AllInstructorInfo(c *gin.Context){
	res := service.AllInstructorService(c)
	c.JSON(200,res)
}
