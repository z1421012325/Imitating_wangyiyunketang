package DB

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm"
	"redigo/redis"

)


var (
	DB *gorm.DB
	RB *redis.Pool
)


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




	rd := &redis.Pool{
				Dial: func() (conn redis.Conn, e error) {
					return redis.Dial("tcp",os.Getenv("REDIS_ADDR"))
				},

				// 最大空闲
				MaxIdle:         10,
				//当为零时，池中的连接数没有限制。
				MaxActive:       0,
				// 超时时间
				IdleTimeout:     5 * time.Second,
				// 当超过最大连接数时直接返回错误 默认false
				Wait:            false,
				// 最大等待时间 默认为0 无限制
				MaxConnLifetime: 0,
			}

	RB = rd
}


