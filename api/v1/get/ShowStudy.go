package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func ShowStudy(c *gin.Context){
	res := service.ShowStudyService(c)
	c.JSON(200,res)
}
