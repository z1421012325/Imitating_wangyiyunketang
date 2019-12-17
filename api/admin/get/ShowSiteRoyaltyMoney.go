package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowSiteRoyaltyMoney(c *gin.Context){

	res := admin.ShowSiteRoyaltyMoneyService()
	c.JSON(200,res)
}
