package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	service2 "demos/service"
	"demos/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ExtractMoneyService struct {
	// 提成用户
	UID int						`json:"uid" form:"uid" binding:"required"`			// GetUserId 得到的uid实际类型为uint64
	// 提成金额
	UserExtractMoney float64	`json:"money" form:"money" binding:"required"`
	// 提取账号
	UserCard string				`json:"card" form:"card" bindging:"required"`
}

type ExtractMoneyData struct {
	Result model.Money  	`json:"result"`
	Uuid   string			`json:"uuid"`
}


func (service *ExtractMoneyService)ExtractMoney(c *gin.Context)*serialize.Response{
	uid := service2.GetUserId(c)

	//fmt.Println(reflect.TypeOf(uid))      // 实际类型为uint64
	//fmt.Println(reflect.TypeOf(service.UID))
	//if service.UID != uid{
	//	return serialize.PswdErr("提取人账号不符合",nil)
	//}
	if service2.CheckUidToUid(service.UID,c){
		return serialize.PswdErr("提取人账号不符合",nil)
	}

	var money model.Money
	DB.DB.Where("u_id = ?",uid).First(&money)
	if money.Money < service.UserExtractMoney || service.UserExtractMoney == 0 {
		return serialize.AccountErr("提取金额大于实际金额或者为0",nil)
	}


	ExtractMoney := service.UserExtractMoney * model.Divide		// 提成金额
	ActualMoney := service.UserExtractMoney - ExtractMoney		// 实际作者获得金额
	SurplusMoney := money.Money - ActualMoney

	//uuid := util.GetUuid()
	uuid := util.RandIntToString()
	UserExtractMoney := fmt.Sprint(service.UserExtractMoney)

	if !util.TransferTransaction(uuid,service.UserCard,UserExtractMoney){
		return serialize.TransactionErr("转账失败",nil)			// todo 定制一个转账错误的信息
	}


	sql1 := "update money set money = ? where money = ? and u_id = ?"		// 当分布式时 可尝试使用消息队列,服务器端生成信息,消费者进行业务处理
	sql2 := "insert into extracts (t_money,divide,actual_money,u_id,number) values (?,?,?,?,?)"

	db1 := DB.DB.Exec(sql1,SurplusMoney,money.Money,money.UID)
	db2 := DB.DB.Exec(sql2,ExtractMoney,model.Divide,ActualMoney,money.UID,uuid)

	ok := DB.Transaction(db1,db2)
	if !ok {
		return serialize.DBErr("数据库转账操作失败",nil)
	}


	// TransferAccounts()  // 未完成 异步开go协程 或者 发消息队列,让消费者来处理

	DB.DB.Where("u_id = ?",uid).First(&money)

	var data ExtractMoneyData
	data.Result = money
	data.Uuid = uuid

	return serialize.Res(data,"")
}



// 转账业务
func TransferAccounts(){}