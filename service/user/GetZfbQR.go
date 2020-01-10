package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"demos/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetZFBQRDate struct {
	WAP 			string		`json:"wap"`
	PC  			string		`json:"pc"`

	Title 			string		`json:"title"`			// 商品标题
	Body  			string		`json:"body"`			// 商品信息
	Price			string		`json:"price"`			// 价格
	OutTradeNo 		string		`json:"out_trade_no"`	// 订单号
}


func GetZFBQR(c *gin.Context) *serialize.Response{

	cid := c.Param("cid")
	var cc model.Curriculums
	DB.DB.Where("c_id = ?",cid).First(&cc)

	OutTradeNo 	:= util.RandIntToString()
	title 		:= "[课程]" + cc.Name
	body 		:= "购买课程"
	price 		:= strconv.FormatFloat(cc.Price,'g', 5, 64)

	returnurl   := c.Request.URL.RequestURI()

	data := &GetZFBQRDate{
		WAP:        util.GetAlipayWapUrl(title,OutTradeNo,price,body,returnurl),
		PC:         util.GetAlipayPageUrl(title,OutTradeNo,price,body,returnurl),
		Title:      title,
		Body:       body,
		Price:      price,
		OutTradeNo: OutTradeNo,
	}


	// 添加订单
	uid := service.GetUserId(c)
	if !addCourseOrder(cid,uid,cc.Price,OutTradeNo){
		return serialize.DBErr("添加订单失败",nil)
	}

	fmt.Println("添加订单完成...")

	return serialize.Res(data,"添加订单成功")
}




func addCourseOrder(cid string,uid interface{},price float64,number string) bool{

	sql1 := "insert into purchases (c_id,u_id,status,price,number) values (?,?,?,?,?)"
	//sql2 := "insert into shopping_carts (c_id,u_id) values (?,?)"

	db1 := DB.DB.Exec(sql1,cid,uid,model.DefaultStatus,price,number)
	//db2 := DB.DB.Exec(sql2,cid,uid)

	if !DB.Transaction(db1){
		return false
	}
	return true
}
