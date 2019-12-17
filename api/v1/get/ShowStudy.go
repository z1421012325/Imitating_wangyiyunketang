package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ShowStudy(c *gin.Context){
	res := user.ShowStudyService(c)
	c.JSON(200,res)
}
