package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowTeacherRoyaltyMoneyDay(c *gin.Context)  {

	res := admin.ShowTeacherRoyaltyMoneyDayService(c)
	c.JSON(200,res)
}