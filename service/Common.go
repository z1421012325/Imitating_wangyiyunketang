package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"demos/DB"
	"demos/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
	得到页码数和每页的个数
	每个http请求的上下文中携带url参数
	参数 : page   -- >>  页码数
		  size   -- >>	每页限制数据个数
 */
func PagingQuery(c *gin.Context)(start,end int){
	var err error
	undetermined1 := c.DefaultQuery("page","0")
	undetermined2 := c.DefaultQuery("size","20")

	page,err := strconv.Atoi(undetermined1)
	if (err != nil || page < 0) {
		page = 0
	}

	size ,err := strconv.Atoi(undetermined2)
	if err != nil {
		size = 20
	}else if (size <=0 ||size >= 50) {
		size = 30
	}

	start = page*size
	// end   = page*size+size
	// end = size
	return start,size
}


/*
	点击热度排行得到页数和每页获取的数据个数 (redis专用)
 */
func ClickRankingPagingQuery(c *gin.Context) (from int64,to int64){
	var err error
	undetermined1 := c.DefaultQuery("from","0")
	undetermined2 := c.DefaultQuery("to","20")

	from , err = strconv.ParseInt(undetermined1,10,64)
	if (err != nil || from < 0) {
		from = 0
	}else if from == 1 {		// 前端传递 1 - 20 redis是从0开始的,所以等于1 就等于从最开始查询
		from = 0
	}

	to , err = strconv.ParseInt(undetermined2,10,64)
	if err != nil {
		to = 20
	}else if (to <=0 ||to >= 50) {
		to = 50
	}

	return from,to - 1			// 查询0 -20 那么表示 0-19的20条数据
}







/*
	验证用户信息 auto
	在http的上下文中找到设置的model数据,模型转换 看是否能验证通过
	返回模型和bool
 */
func VerifyUser(c *gin.Context)(model.User,bool){
	s := sessions.Default(c)
	user := s.Get("user")
	if user != nil{
		if chechuser, ok := user.(model.User); ok {

			// pass 掉 session中取出的值并不对等

			// 再次验证 user_id 是否存在和相等
			//uid := s.Get("user_id")
			//if string(chechuser.ID) == uid {
			//	return chechuser,true
			//}

			return chechuser,true
		}
	}
	return model.User{},false
}



/*
	取得session中的用户id
	在上下文中找到 设置(登录时设置)的值
	注意 经过测试这里返回的user_id 值类型为uint64
 */
func GetUserId(c *gin.Context) interface{}{
	s := sessions.Default(c)
	return s.Get("user_id")
}


/*
	取得session中的用户adminid
	在上下文中找到 设置(登录时设置)的值
	注意 经过测试这里返回的adminid 值类型为uint64
*/
func GetAdminId(c *gin.Context) interface{}{
	s := sessions.Default(c)
	return s.Get("admin_id")
}


/*
	根据结构体(表单提交数据)和上下文中的值进行对比
 */
func CheckUidToUid(checkint int,c *gin.Context) bool{
	uid := GetUserId(c)			// 得到uint64 位 进行对比
	if uid != uint64(checkint){
		return false
	}
	return true
}








// -------------------------------------  cache redis 缓存相关function






/*
	获得 redis 中string类型的数据的值
	返回string
 */
func GetCacheTypeStr(keys []string) string {
	key := strings.Join(keys,"_")		 // 在redis中做一个切割 xxx_xxx_xxx 之类的 不容易混淆
	return DB.GetCache(key)
}

/*
	在redis中 以string数据为key,设置带有时间限制的数据
 */
func SetCacheTypeStr(keys []string,value interface{},tll time.Duration){
	key := strings.Join(keys,"_")
	v,_ := json.Marshal(value)

	if tll <= 0 {
		tll = DB.CacheTllLevel2					// 暂时缓存时间为 5 min
	}

	DB.SetCache(key,string(v),tll)
}

/*
	删除缓存
 */
func DelCacheTypeStr(){}






/*
	有序集合添加数据
 */
func SetCacheTypeAgg(ordername string,key string){
	DB.ZaddList(ordername,key)
}


/*
	获得 redis 中有序集合类型的数据的值
	返回 float64 位 ,但是由乱码 需要进行string转换
*/
func GetCacheTypeAggReturnFloat(ordername string,key string) float64{
	return DB.Zscore(ordername, key)
}

func GetCacheTypeAggReturnString(ordername string,key string) string{
	return fmt.Sprint(DB.Zscore(ordername, key))
}


/*
	根据start和end 返回这个区间之内的数据   返回的是key  不会返回incr值
 */
func GetCacheTypeAggScope(ordername string,start,end int64)  []string {
	return DB.Zrevrange(ordername,start,end)
}

/*
	增加在有序集合中key的增量值 +1
*/
func SetOrderIncr(ordername string,key string){
	DB.ZIncr(ordername,key,1)
}



/*
	删除有序集合中的数据
 */
func DelOrderData(ordername string,key string) bool {
	if !DB.DelOrderlyelement(ordername,key){
		return false
	}
	return true
}






/*
	在hash中添加数据
 */
func SetHashData(hashname string,key string,value interface{}){
	field := make(map[string]interface{})
	field[key] = value

	DB.AddCacheHash(hashname,field)

}


/*
	在hash中获得数据
 */
func GetHashData(hashname ,field string) string{
	return DB.GetCacheHash(hashname,field)
}


/*
	在hash中删除数据
 */
func DelHashData(hashname,field string){
	DB.DelCacheHash(hashname,field)
}








