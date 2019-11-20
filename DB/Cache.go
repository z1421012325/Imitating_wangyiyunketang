package DB

import (
	"github.com/go-redis/redis"
	"log"
	"strings"
	"time"
)

// 查询cache
func GetCache(str []string) string{
	key := strings.Join(str,"_")

	values , err := RB.Get(key).Result()
	if err != nil{
		return values		// nil
	}

	return values
}

// 添加cache
func SetCache(key , value string,TLL time.Duration) {
	err := RB.Set(key,value,TLL).Err()
	if err != nil {
		log.Print(err)   // todo 日志异常一下 消息队列
	}
}
// 添加有序集合
func ZaddList(key string,member string)bool{
	err := RB.ZAdd(key,redis.Z{
		Score:  0,
		Member: member,
	}).Err()
	if err != nil{
		log.Print(err)
		return false
	}
	return true
}



// 增加一个int类型的值 通常用在排行榜,观看人数,或者频繁操作数据库的信息,但是不要忘记redis要持久化操作
func Incr(key string, value int64) bool{
	err := RB.IncrBy(key,value).Err()
	if err != nil {
		log.Print(err)			// 日志异常一下 消息队列
		return false
	}
	return true
}
func ZIncr(key string, increment float64, member string)bool{
	err := RB.ZIncrBy(key,increment,member).Err()
	if err != nil {
		log.Print(err)			// 日志异常一下 消息队列
		return false
	}
	return true
}

// 获取有序集合的成员数
func Count(key string) int64{
	count,_ := RB.ZCard(key).Result()
	return count
}

// 移除有序集中的成员
func DelOrderlyelement(key string, members ...interface{})bool{
	err := RB.ZRem(key,members).Err()
	if err != nil {
		log.Print(err)		//
		return false
	}
	return true
}

// 返回指定区间数据
func Zrange(key string,start,end int64) []string{
	// limit start end  or start size ?
	values, _ := RB.ZRange(key,start,end).Result()
	return values
}


// 返回有序集中，成员的分数值
func Zscore(key string, members string) (score float64) {
	score,_ =RB.ZScore(key,members).Result()
	return
}


// 返回有序集中成员的排名  最大值为0  所以+1
func ZrevRank(key,member string) int64{
	index , _ := RB.ZRevRank(key,member).Result()
	return index+1
}

// 返回有序集中，指定区间内的成员
func Zrevrange(key string,start,end int64) []string {
	value ,_ :=RB.ZRevRange(key,start,end).Result()
	return value
}






















// lock加锁操作
func SetLockWriter(key,value string,TLL time.Duration){
	RB.SetNX(key,value,TLL)
}
func SetZLockWriter(key,value string,TLL time.Duration){

}

// 删除加锁
func DelLock(){}
func ZDelLock(){}

// 查询是否为加锁数据
func QueryLock(){}
func ZQueryLock(){}