package main

import (
	"fmt"
	"os"

	"demos/server"
)


// 开启服务
func main(){
	fmt.Println("run!!!")
	fmt.Println(os.Getenv("MYSQL_HOST"))

	server := server.NewRouter()
	_ = server.Run(os.Getenv("SERVER_PORT"))


}
