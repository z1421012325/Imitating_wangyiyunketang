package get

import (
	"github.com/gin-gonic/gin"
	"demos/service/admin"
	)

func ShowVideoTotal(c *gin.Context){

	res := admin.ShowTotalService()
	c.JSON(200,res)
}


