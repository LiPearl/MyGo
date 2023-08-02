package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	// 将 []byte转换成16进制
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 获取当前时间字符串
func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func getTimeInt() int64 {
	return time.Now().Unix()
}

func main() {
	str := "abcdef"
	fmt.Printf("MD5(%s):%s\n", str, MD5("abcdef"))
	fmt.Printf("current time str : %s\n", getTimeStr())
	fmt.Printf("current time unix : %d\n", getTimeInt())

}
