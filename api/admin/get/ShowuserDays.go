package get

import (
	"github.com/gin-gonic/gin"
	"demos/service/admin"
)

func ShowUserDays(c *gin.Context){

	res := admin.ShowUserDaysService(c)
	c.JSON(200,res)
}
