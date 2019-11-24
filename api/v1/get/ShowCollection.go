package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func ShowCollection(c *gin.Context){
	res := service.ShowCollectionService(c)
	c.JSON(200,res)
}
