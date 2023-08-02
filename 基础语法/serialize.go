package main

import (
	"encoding/json"
	"fmt"
)

// 这个结构体中使用了json:"code"和json:"msg"的标签，
// 用于指定在JSON序列化时各字段对应的名称。
// 例如，在对该结构体进行JSON编码时，Code会被编码为"code"，Message会被编码为"msg"。这种方式可以帮助我们将Go语言中的结构体转换为JSON格式，并且在网络传输过程中方便地进行解码和处理。
type Result struct {
	Code    int    `json:code`
	Message string `json:msg`
}

func main() {
	var res1 Result
	res1.Code = 200
	res1.Message = "success"
	jsons, errs := json.Marshal(res1)
	if errs != nil {
		fmt.Println("json marsha1 error:", errs)
	}

	// 由于变量jsons是一个byte类型的切片，它在输出时会自动转换成字符串格式，并显示为一组十六进制数字序列
	//json data: [123 34 67 111 100 101 34 58 50 48 48 44 34 77 101 115 115 97 103 101 34 58 34 115 117 99 99 101 115 115 34 125]
	// fmt.Println("json data:", jsons)
	fmt.Println("json data:", string(jsons))

	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarsha1 error:", errs)
	}
	fmt.Println("res2", res2)
}
