package main

import (

	_"demos/conf"
	"demos/server"
	"os"

)


// 开启服务
func main(){

	server := server.NewRouter()
	_ = server.Run(os.Getenv("SERVER_PORT"))

}
