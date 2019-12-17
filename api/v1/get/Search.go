package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context){
	res := user.SearchService(c)
	c.JSON(200,res)
}
