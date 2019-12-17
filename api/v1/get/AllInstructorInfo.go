package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func AllInstructorInfo(c *gin.Context){
	res := user.AllInstructorService(c)
	c.JSON(200,res)
}
