package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Student(c *gin.Context){
	res := user.StudentService(c)
	c.JSON(200,res)
}
