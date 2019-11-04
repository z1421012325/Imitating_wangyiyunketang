package main

import (
	"fmt"
	_ "demos/conf"
	"os"
)


// 开启服务
func main(){
	fmt.Println("run!!!")

	fmt.Println(os.Getenv("MYSQL_HOST"))

}
