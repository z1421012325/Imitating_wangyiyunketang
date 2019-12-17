package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func TeacherCurriculumVideo(c *gin.Context){
	res := user.TeacherCurriculumVideoService(c)
	c.JSON(200,res)
}
