package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)


// 课程页面
func Introduction(c *gin.Context){
	res := service.IntroductionService(c)
	c.JSON(200,res)
}
