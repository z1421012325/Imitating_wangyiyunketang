package get

import (
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func ShowListVideo(c *gin.Context){

	res := admin.ShowListVideoService(c)
	c.JSON(200,res)
}
