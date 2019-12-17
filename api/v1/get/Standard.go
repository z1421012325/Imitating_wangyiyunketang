package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Standard(c *gin.Context){
	res := user.StandardService(c)
	c.JSON(200,res)
}
