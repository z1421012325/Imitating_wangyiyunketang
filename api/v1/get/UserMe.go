package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func UserMe(c *gin.Context){
	res := service.UserMeService(c)
	c.JSON(200,res)
}
