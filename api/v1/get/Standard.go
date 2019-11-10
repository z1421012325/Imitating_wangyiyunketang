package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Standard(c *gin.Context){
	res := service.StandardService(c)
	c.JSON(200,res)
}
