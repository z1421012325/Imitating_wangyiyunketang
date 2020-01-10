package conf

import (
	"godotenv"
	"os"
)


// 加载两次文件,如果文件不加载,尝试在系统环境中查看是否含有变量
func init(){
	err := godotenv.Load("./config.env")
	if err != nil {

		err = godotenv.Load("config.env")
		if err != nil {
			ok := os.Getenv("IS_LOAD_CONFIG_FILE")
			if ok != "true"{
				panic(err.Error())
			}
		}
	}
}