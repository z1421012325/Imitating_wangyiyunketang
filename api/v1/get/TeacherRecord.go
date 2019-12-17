package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func TeacherRecord(c *gin.Context){
	res := user.TeacherRecordService(c)
	c.JSON(200,res)
}
