package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Catalog(c *gin.Context){
	res := user.CatalogService(c)
	c.JSON(200,res)
}