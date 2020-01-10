package post

import (
	"demos/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ZfbCallback(c *gin.Context){

	//fmt.Println(c.Request.Body)

	//util.VerifyNotify(c)

	fmt.Println("支付宝回调函数 启用~")

	user.ZfbCallbackService(c)


}