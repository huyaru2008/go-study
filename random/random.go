package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//随机字符串
func RandStr(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 随机数验证码
func GenValidateCode(len int) string {
	numbers := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < len; i++ {
		fmt.Fprintf(&sb, "%d", numbers[rand.Intn(10)])
	}
	return sb.String()
}
