package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Record(c *gin.Context){
	res := user.RecordService(c)
	c.JSON(200,res)
}
