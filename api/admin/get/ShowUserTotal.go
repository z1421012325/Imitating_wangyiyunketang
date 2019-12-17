package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowUserTotal(c *gin.Context){

	res := admin.ShowUserTotalService()
	c.JSON(200,res)
}
