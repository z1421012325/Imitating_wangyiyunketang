package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"

)

func GetZFBQR(c *gin.Context){
	res := user.GetZFBQR(c)
	c.JSON(200,res)
}
