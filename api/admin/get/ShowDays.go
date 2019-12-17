package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowCurriculumDays(c *gin.Context){

	res := admin.ShowDaysService(c)
	c.JSON(200,res)
}
