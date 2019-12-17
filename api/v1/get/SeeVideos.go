package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func SeeVideos(c *gin.Context){
	res := user.SeeVideosService(c)
	c.JSON(200,res)
}
