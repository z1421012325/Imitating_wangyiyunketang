package util

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/satori/go.uuid"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func RandIntToString() string {
	var letterRunes = []rune("1234567890")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}






// 得到一个uuid
func GetUuid() string {
	u1 := uuid.Must(uuid.NewV4(),nil)
	return u1.String()
}

func CheckUuid(u1 string) bool {
	_, err := uuid.FromString(u1)
	if err != nil {
		return false
	}
	return true
}



// 得到本机ip
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}


