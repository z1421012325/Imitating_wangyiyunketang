package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ShowListClick(c *gin.Context){
	res := user.ShowListClickService(c)
	c.JSON(200,res)
}