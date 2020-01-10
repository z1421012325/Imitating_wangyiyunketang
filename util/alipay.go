package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	alipay2 "github.com/smartwalle/alipay/v3"
)


/*
https://github.com/smartwalle/alipay 支付宝接口 文档
 */

const (
	alipayWAPProductCode = "QUICK_WAP_WAY"						//wap url
	alipayPCProductCode  = "FAST_INSTANT_TRADE_PAY"				//pc url

	// ps
	WAIT_BUYER_PAY 		 = "WAIT_BUYER_PAY"      				//（交易创建，等待买家付款）
	TRADE_CLOSED		 = "TRADE_CLOSED"						//（未付款交易超时关闭，或支付完成后全额退款）
	TRADE_SUCCESS		 = "TRADE_SUCCESS"						//（交易支付成功）
	TRADE_FINISHED		 = "TRADE_FINISHED"						//（交易结束，不可退款）

	PAYEE_TYPE 			 = "ALIPAY_LOGONID"						// 转账账号类型 支付宝
	//PAYEE_ACCOUNT		 = "bpsmve6191@sandbox.com"				// 支付宝账号
	PAYER_SHOW_NAME		 = "xxxx公司"							// 转账人姓名

	ALIPAY_ACCOUNT		 = "bpsmve6191@sandbox.com"				// 支付宝沙盒账号 支付密码默认为111111
)

var (
	appid 				= os.Getenv("app_id")					// 注册app的id
	appprivate 			= os.Getenv("private_Key")				// app进行信息加密的私钥
	alipaypublic 		= os.Getenv("public_key")				// 支付宝在提交公钥之后给与的支付宝公钥 用来解密支付宝传递过来的信息
	isProduction 		= false										// 沙盒账号则为false 真是账号为 true

	NotifyURL 			= "http://" + GetLocalIP() + os.Getenv("SERVER_PORT") + "/zfb/callback"

	Client 				*alipay2.Client								// alipay单例
)


/*
	初始化alipay.client,单例
 */
func init()  {
	// 沙盒账号则为false 真是账号为 true
	if os.Getenv("GIN_MODE") != "debug"{
		isProduction = true
	}

	client ,err := alipay2.New(appid,appprivate,isProduction)
	_ = client.LoadAliPayPublicKey(alipaypublic)						// 加载支付宝公钥
	if err != nil {
		fmt.Println(err)
	}
	Client = client
}




/*
	返回跳转支付宝支付跳转url   app wap 手机网页版
	参数:Subject -- >>	标题
		OutTradeNo -- >> 订单号
		TotalAmount -- >> 金额
		body -- >> 其他信息
		ReturnURL -- >> 当完成支付之后进行跳转的地址(url) 可选
	返回 : 转账跳转url

 */
func GetAlipayWapUrl(Subject,OutTradeNo,TotalAmount,body string,ReturnURL... string) string {

	var p = alipay2.TradeWapPay{}
	p.NotifyURL 	= NotifyURL

	p.ReturnURL 	= ReturnURL[0]

	p.Subject 		= Subject
	p.OutTradeNo 	= OutTradeNo
	p.TotalAmount 	= TotalAmount
	p.Body 			= body
	p.ProductCode 	= alipayWAPProductCode

	url, err := Client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	return payURL
}




/*
	返回跳转支付宝支付跳转url   pc版
	参数:Subject -- >>	标题
		OutTradeNo -- >> 订单号
		TotalAmount -- >> 金额
		body -- >> 其他信息
		ReturnURL -- >> 当完成支付之后进行跳转的地址(url) 可选
	返回 : 转账跳转url

 */
func GetAlipayPageUrl(Subject,OutTradeNo,TotalAmount,body string,ReturnURL... string) string {

	var p = alipay2.TradePagePay{}
	p.NotifyURL 	= NotifyURL
	p.ReturnURL 	= ReturnURL[0]									// 支付完成之后跳转的url  可选

	p.Subject 		= Subject
	p.OutTradeNo 	= OutTradeNo
	p.TotalAmount	= TotalAmount
	p.Body 			= body
	p.ProductCode 	= alipayPCProductCode

	url, err := Client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	return payURL
}



/*
	查询订单信息 支付完成则返回true
	参数 : OutTradeNo -- >>  订单号
	返回 : 转账确定

	ps 查询的时支付宝支付信息,但是商户数据库暂时没有更改订单状态
 */
func QueryOutTradeNo(OutTradeNo string)  bool {

	var p = alipay2.TradeQuery{}
	p.OutTradeNo = OutTradeNo

	res,err := Client.TradeQuery(p)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if !res.IsSuccess(){
		return false
	}
	return true
}


/*
	异步验证通知 返回支付宝确定通知 返回通知是否支付和商户订单号
	接受alipay的中设置的回调(post请求) 得到其中传递过来的参数,如果能映射alipay自定义的struct并sign通过验证 则表示正确
	参数: *gin.content(request)    -- >> http 的上下文请求,gin给打包一层
	返回 : 订单号

	ps >> gin框架是包裹net/http
 */
func VerifyNotify (c *gin.Context) (string) {

	var noti, _ = Client.GetTradeNotification(c.Request)
	if noti != nil {
		fmt.Println("支付成功")
	} else {
		fmt.Println("支付失败")
		return ""
	}
	alipay2.AckNotification(c.Writer)

	return noti.OutTradeNo
}















/*
	转账业务 -- >> 个人支付宝
	参数 : OutBizNo			-- >> 订单号
			Account			-- >> 收款账号
			Amount			-- >> 金额
			Remark			-- >> 备注 可选
	返回 : 转账确定
 */
func TransferTransaction(OutBizNo string,Account string,Amount string,Remark... string ) bool {

	var req alipay2.FundTransToAccountTransfer
	req.OutBizNo = OutBizNo										// 商户转账唯一订单号
	req.PayeeType = PAYEE_TYPE									// 收款方账户类型,"ALIPAY_LOGONID":支付宝帐号
	req.PayeeAccount = Account									// 收款方账户
	req.Amount = Amount											// 转账金额
	req.PayerShowName = PAYER_SHOW_NAME							// 付款方显示姓名
	req.Remark = Remark[0]										// 转账备注


	result,err := Client.FundTransToAccountTransfer(req)
	if err != nil {
		fmt.Println("转账异常")
		return false
	}

	if !result.IsSuccess(){
		fmt.Println("转账失败")
		return false
	}
	fmt.Println("转账成功")
	return true
}



/*
	查询个人转账交易状态
 */
func queryTransferTransaction(){
	//var req alipay2.FundTransOrderQuery
	// Client.FundTransOrderQuery(req)
}






