package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func SearchTag(c *gin.Context){
	res := service.SearchTagService(c)
	c.JSON(200,res)
}
