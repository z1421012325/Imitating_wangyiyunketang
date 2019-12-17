package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowTotalMoney(c *gin.Context){

	res := admin.ShowTotalMoneyService()
	c.JSON(200,res)
}
