package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func NowShowlist(c *gin.Context){
	res := user.NowCurriculumService(c)
	c.JSON(200,res)
}
