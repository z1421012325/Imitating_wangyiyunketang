package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowDayMoney(c *gin.Context)  {
	res := admin.ShowDayMoneyService(c)
	c.JSON(200,res)
}
