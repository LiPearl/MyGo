package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"time"
)

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 获取当前时间字符串
func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func getTimeInt() int64 {
	return time.Now().Unix()
}

// 生成签名
func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		for i := 0; i < len(key); i++ {
			if i==0 {
				str= fmt.Sprintf("%v=%v",key[i],params[key[i]])
			}
		else{
			
		}
	}
}
func main() {
	str := "12345"
	fmt.Printf("md5[%s]:%s", str, MD5(str))
	fmt.Printf("current time str :%s\n", getTimeStr())
	fmt.Printf("current time unix :%d\n", getTimeInt())
}
