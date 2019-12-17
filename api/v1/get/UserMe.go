package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func UserMe(c *gin.Context){
	res := user.UserMeService(c)
	c.JSON(200,res)
}
