package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Student(c *gin.Context){
	res := service.StudentService(c)
	c.JSON(200,res)
}
