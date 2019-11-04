package main

import (
	_ "demos/conf"
	"fmt"
	"os"

	"demos/server"
)


// 开启服务
func main(){
	fmt.Println("run!!!")
	fmt.Println(os.Getenv("MYSQL_HOST"))

	server := server.NewRouter()
	server.Run(os.Getenv("SERVER_PORT"))

}
