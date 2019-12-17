package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowMonthMoney(c *gin.Context){

	res := admin.ShowMonthMoneyService(c)
	c.JSON(200,res)
}
