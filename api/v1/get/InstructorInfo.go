package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func InstructorInfo(c *gin.Context){
	res := user.InstructorInfoService(c)
	c.JSON(200,res)
}
