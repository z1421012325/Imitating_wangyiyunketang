package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func SearchTag(c *gin.Context){
	res := user.SearchTagService(c)
	c.JSON(200,res)
}
