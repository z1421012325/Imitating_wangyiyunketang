package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func HaveMoney(c *gin.Context){
	res := user.HaveMoneyService(c)
	c.JSON(200,res)
}
