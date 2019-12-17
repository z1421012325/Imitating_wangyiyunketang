package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ShowCollection(c *gin.Context){
	res := user.ShowCollectionService(c)
	c.JSON(200,res)
}
