package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowProhibitUsers(c *gin.Context){
	res := admin.ShowProhibitUsersService(c)
	c.JSON(200,res)
}
