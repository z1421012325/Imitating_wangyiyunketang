package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowSiteRoyaltyMoneyDay(c *gin.Context){

	res := admin.ShowSiteRoyaltyMoneyDayService(c)
	c.JSON(200,res)
}
