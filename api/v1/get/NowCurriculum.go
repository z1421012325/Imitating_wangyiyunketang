package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func NowCurriculum(c *gin.Context){
	res := service.NowCurriculumService(c)
	c.JSON(200,res)
}
