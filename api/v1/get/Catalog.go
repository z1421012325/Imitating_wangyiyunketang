package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Catalog(c *gin.Context){
	res := service.CatalogService(c)
	c.JSON(200,res)
}