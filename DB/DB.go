package DB

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm"
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



/*
 	mysql 和redis 连接池连接初始化
 */
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

	_,err = rb.Ping().Result()
	if err != nil {
		panic(err)
	}

	RB = rb
}



