package main

import (
	"os"

	_"demos/conf"

	"demos/server"
)


// 开启服务
func main(){

	server := server.NewRouter()
	_ = server.Run(os.Getenv("SERVER_PORT"))

}
