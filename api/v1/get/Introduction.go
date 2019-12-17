package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)


// 课程页面
func Introduction(c *gin.Context){
	res := user.IntroductionService(c)
	c.JSON(200,res)
}
