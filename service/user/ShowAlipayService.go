package user

import (
	"demos/DB"
	"demos/model"
	"demos/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ZfbCallbackService(c *gin.Context){

	outtradeon := util.VerifyNotify(c)

	sql1 := "update purchases set status = ? where number = ?"
	db1  := DB.DB.Exec(sql1,model.CompleteStatus,outtradeon)

	if !DB.Transaction(db1){
		fmt.Println("订单保存失败")
	}
	fmt.Println("订单保存成功")
	return
}
