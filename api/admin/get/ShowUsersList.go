package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowUsersList(c *gin.Context){

	res := admin.ShowUsersListService(c)
	c.JSON(200,res)
}
