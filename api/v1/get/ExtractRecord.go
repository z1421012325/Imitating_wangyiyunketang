package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ExtractRecord(c *gin.Context){
	res := user.ExtractRecordService(c)
	c.JSON(200,res)
}