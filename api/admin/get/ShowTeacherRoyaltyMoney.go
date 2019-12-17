package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowTeacherRoyaltyMoney(c *gin.Context){

	res := admin.ShowTeacherRoyaltyMoneyService()
	c.JSON(200,res)
}
