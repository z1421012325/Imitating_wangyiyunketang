package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"demos/util"
	"github.com/gin-gonic/gin"
)

type QueryZfbIssuccessService struct {
	OutTradeNo 	string				`json:"out_trade_no" form:"out_trade_no" binding:"required"`
	CID 		int					`json:"cid" form:"cid" binging:"required"`
}




func (service *QueryZfbIssuccessService) QueryZfbIssuccess(c *gin.Context) *serialize.Response{

	// 支付宝查询是否支付成功
	if !util.QueryOutTradeNo(service.OutTradeNo){
		return serialize.Res(nil,"未完成支付")
	}

	/*
	modify
	code :
		// 成功则在数据库中修改支付状态(下单时直接在数据库中生成,但是状态为未支付),修改成功 返回查询的支付信息
		//uid := service2.GetUserId(c)
		//if !service.modifyCourseOrder(service.OutTradeNo,uid){
		//	return serialize.DBErr("订单异常,请联系客服",nil)
		//}
	msg  : 想了想 还是不在查询支付是否成功接口来修改订单状态,修改订单状态由支付宝异步callback 设置的url来修改状态
			这个接口只用来设置 购买记录(非支付记录)
	 */

	uid := service2.GetUserId(c)
	service.modifyCourseOrder(service.OutTradeNo,uid)

	// 返回
	return serialize.Res(nil,"支付成功")

}



func (service *QueryZfbIssuccessService)modifyCourseOrder(number string,uid interface{}) bool{

	//sql1 := "update purchases set status = ? where number = ? and u_id = ?"
	//db1  := DB.DB.Exec(sql1,model.CompleteStatus,number,uid)

	sql2 := "insert into shopping_carts (c_id,u_id) values (?,?)"
	db2  := DB.DB.Exec(sql2,service.CID,uid)

	if !DB.Transaction(db2){
		return false
	}

	return true
}