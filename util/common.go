package util

import (
	"math/rand"
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





func GetUuid()uuid.UUID{
	u1 := uuid.Must(uuid.NewV4(),nil)
	return u1
}

func CheckUuid(u1 string)bool{
	_, err := uuid.FromString(u1)
	if err != nil {
		return false
	}
	return true
}



