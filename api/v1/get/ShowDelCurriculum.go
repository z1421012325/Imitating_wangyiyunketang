package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ShowDelCurriculum(c *gin.Context){
	res := user.ShowDelCurriculumService(c)
	c.JSON(200,res)
}
