package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context){
	res := service.SearchService(c)
	c.JSON(200,res)
}
