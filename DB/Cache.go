package DB

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)


const (
	CacheTimeL1 = time.Minute * 1			// 缓存时间
	CacheTimeL2 = time.Minute * 5
	CacheTimeL3 = time.Minute * 10
	CacheTimeL4 = time.Minute * 30
	CacheTimeL5 = time.Minute * 60
)



// --------------------------string ---------------------------------------------

/*
	根据字段在redis中查找 string 对应的数据类型
	例如: set xxx 123      --- >> 设置string类型的数据 			key > xxx  value > 123
	获取: get xxx     	   --- >> 返回 value
	参数: str  >> 多个string
 */
func GetCache(key string) string{

	values , err := RB.Get(key).Result()
	if err != nil{
		return values		// nil
	}
	return values
}


/*
	在redis中添加字段 类型为string
	例如: set zxc 123      --- >>           key > zxc   value > 123
	参数: key ------- string类型的key
		 value -------- 传递进来key的value值
	     tll -------- 过期时间
 */
func SetCache(key , value string,TLL time.Duration) {
	err := RB.Set(key,value,TLL).Err()
	if err != nil {
		log.Print(err)   // todo 日志异常一下 消息队列
	}
}









// -------------------------------set ------------------------------------------

/*
	获取有序集合中的数据
	参数: ordername ------- 	有序集合的名称
		 member -------- 有序集合中设置的键 key
 */
func GetAgg(){}




/*
	添加有序集合   默认value 为 0
	参数: ordername ------- 	有序集合的名称
		 member -------- 有序集合中设置的键 key
 */
func ZaddList(ordername string,member string) bool {
	err := RB.ZAdd(ordername,redis.Z{
		Score:  0,
		Member: member,
	}).Err()
	if err != nil{
		log.Print(err)
		return false
	}
	return true
}



/*
	在无序集合中增加集合的Score  默认每次+1
	参数: ordername ------- 	集合的名称
		 member -------- 有序集合中设置的键 key
	ps 增加一个int类型的值 通常用在排行榜,观看人数,或者频繁操作数据库的信息,但是不要忘记redis要持久化操作
 */
//func Incr(ordername string, incr int64) bool {
//	err := RB.IncrBy(ordername,incr).Err()
//	if err != nil {
//		log.Print(err)			// 日志异常一下 消息队列
//		return false
//	}
//	return true
//}



/*
	有序集合中数据的score增加 默认+1
	参数: ordername ------- 	有序集合的名称
		 member -------- 有序集合中设置的键 key
		 increment -------- 有序集合中键 key 增加的值

	例如: zadd runoob(有序集合名称) 1(增量) rabitmq(有序集合中的key)
 */
func ZIncr(ordername,member string,increment float64) bool {
	err := RB.ZIncrBy(ordername,increment,member).Err()
	if err != nil {
		log.Print(err)			// 日志异常一下 消息队列
		return false
	}
	return true
}

/*
	获取有序集合的成员数
	参数: key  ---- >> 有序集合中有效的总个数

	返回 个数(int类型)
 */
func Count(key string) int64{
	count,_ := RB.ZCard(key).Result()
	return count
}



/*
	移除有序集中的成员
	参数: ordername  ---- >> 有序集合的名字
		 members	---- >> 需要在有序集合中删除的key

 */
func DelOrderlyelement(ordername string, members ...interface{})bool{
	err := RB.ZRem(ordername,members).Err()
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}



///*
//	返回指定区间数据 非排序再进行
//	参数: ordername  ---- >> 有序集合的名字
//		 start	---- >> 需求区间的起点
//		 end	---- >> 需求区间的终点
//	返回: 在查找到的有序集合区间 key 值
//
//	ps : limit start end  or start size ?
// */
//func Zrange(ordername string,start,end int64) []string{
//	values, _ := RB.ZRange(ordername,start,end).Result()
//	return values
//}


/*
	返回有序合集中单个(key)数值的增量值
	参数: ordername  ---- >> 有序集合的名字
		 members	---- >> 需求区间的起点
 */
func Zscore(ordername string, members string) (score float64) {
	score,_ =RB.ZScore(ordername,members).Result()
	return
}


/*
	返回有序集中成员的排名(默认为排序之后的数据)  索引默认为0  所以+1
	参数: ordername   --- >> 有序集合的name
		 member      --- >> 有序集合中的key
	返回 : 返回member在有序集合中的index
 */
func ZrevRank(key,member string) int64{
	index , _ := RB.ZRevRank(key,member).Result()
	return index+1
}

/*
	返回有序集中指定区间内的成员(默认为经过排序sort)
	参数 : ordername     ---- >>> 有序集合名字
		  stard			---- >>> 排序之后开始的位置
		  end			---- >>> 排序之后区间的末尾

 */
func Zrevrange(ordername string,start,end int64) []string {
	value ,_ :=RB.ZRevRange(ordername,start,end).Result()
	return value
}





// ---------------------------------------list ----------------------------

// -------------------------------------hash ----------------------------


/*
	在hash(字典,map)中添加一个,如果拥有则覆盖,没有则创建
	参数 : hashname  --- >> hash 在redis中存在的name
		  field    --- >>  保存在hashname中 key 与 value
 */
func AddCacheHash(hashname string,field map[string]interface{}) {
	RB.HMSet(hashname,field)
}

/*
	查询hash中的数据
 */
func GetCacheHash(hashname ,field string) string {
	result,_ := RB.HGet(hashname,field).Result()
	return result
}


/*
	删除hash中的数据
 */
func DelCacheHash(hashname,field string)  {
	 RB.HDel(hashname,field)
}








// ------------------------- 分布式锁 操作---------------------

// lock加锁操作
func SetLockWriter(key,value string,TLL time.Duration){
	RB.SetNX(key,value,TLL)
}
func SetZLockWriter(key,value string,TLL time.Duration){}

// 删除加锁
func DelLock(){}
func ZDelLock(){}

// 查询是否为加锁数据
func QueryLock(){}
func ZQueryLock(){}