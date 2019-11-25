package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func ShowDelCurriculum(c *gin.Context){
	res := service.ShowDelCurriculumService(c)
	c.JSON(200,res)
}
