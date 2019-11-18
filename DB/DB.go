package DB

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm"
	"github.com/go-redis/redis"

)

var (
	DB *gorm.DB
	// RB *redis.Pool
	RB *redis.Client

	CacheTllLevel1 time.Duration = 60 * 1  *  time.Second
	CacheTllLevel2 time.Duration = 60 * 2  *  time.Second
	CacheTllLevel3 time.Duration = 60 * 4  *  time.Second
	CacheTllLevel4 time.Duration = 60 * 8  *  time.Second
	CacheTllLevel5 time.Duration = 60 * 16 *  time.Second
	CacheTllLevel6 time.Duration = 60 * 32 *  time.Second
)

//const (
//	CacheTllLevel1 int = 60
//	CacheTllLevel2 int = 60 * 2
//	CacheTllLevel3 int = 60 * 4
//	CacheTllLevel4 int = 60 * 8
//	CacheTllLevel5 int = 60 * 16
//	CacheTllLevel6 int = 60 * 32
//)


func init(){
	name := os.Getenv("MYSQL_NAME")
	pswd := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")

	db,err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",name,pswd,host,port,dbname))
	if err != nil {
		panic(err.Error())
	}

	// 开发模式和生产模式日志开关,环境
	debug := os.Getenv("GIN_MODE")
	if gin.DebugMode != debug {
		db.LogMode(false)
	}else {
		db.LogMode(true)
	}

	// 设置可重用连接的最大时间量 如果d<=0，则永远重用连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//设置到数据库的最大打开连接数 如果n<=0，则不限制打开的连接数 默认值为0
	db.DB().SetMaxOpenConns(0)
	// 设置空闲中的最大连接数 默认最大空闲连接数当前为2 如果n<=0，则不保留空闲连接
	db.DB().SetMaxIdleConns(10)
	DB = db




	//rd := &redis.Pool{
	//			Dial: func() (conn redis.Conn, e error) {
	//				return redis.Dial("tcp",os.Getenv("REDIS_ADDR"))
	//			},
	//
	//			// 最大空闲
	//			MaxIdle:         10,
	//			//当为零时，池中的连接数没有限制。
	//			MaxActive:       0,
	//			// 超时时间
	//			IdleTimeout:     5 * time.Second,
	//			// 当超过最大连接数时直接返回错误 默认false
	//			Wait:            false,
	//			// 最大等待时间 默认为0 无限制
	//			MaxConnLifetime: 0,
	//		}
	//
	//RB = rd
	rb := redis.NewClient(&redis.Options{
		Addr:               os.Getenv("REDIS_ADDR"),
		Password:           "",
		DB:                 0,
		DialTimeout:        0,
		ReadTimeout:        time.Second,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       10,
		MaxConnAge:         0,
	})

	//_,err = rb.Ping().Result()
	//if err != nil {
	//	panic(err)
	//}
	RB = rb
}





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


























// louk加锁操作
func SetLoukWriter(key,value string,TLL time.Duration){
	RB.SetNX(key,value,TLL)
}
func SetZLoukWriter(key,value string,TLL time.Duration){

}

// 删除加锁
func DelLouk(){}
func ZDelLouk(){}

// 查询是否为加锁数据
func QueryLouk(){}
func ZQueryLouk(){}


