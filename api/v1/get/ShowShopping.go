package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ShowShopping(c *gin.Context){
	res := user.ShowShoppingService(c)
	c.JSON(200,res)
}
