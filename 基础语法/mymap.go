package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p1 map[int]string
	// p1[0] = "test0"  未初始化（分配内存），只定义了一个变量
	// 直接赋值，会报panic错位
	p1 = make(map[int]string)
	p1[1] = "test1"
	fmt.Println("p1:", p1)

	var p2 = map[int]string{1: "test1"}
	fmt.Println("p2:", p2)

	p3 := map[int]string{1: "test1"}
	fmt.Println("p3:", p3)

	var p4 map[int]string = make(map[int]string)
	p4[1] = "test1"
	fmt.Println("p4:", p4)

	p5 := map[int]string{}
	p5[1] = "test1"
	fmt.Println("p5:", p5)

	//换行初始化必须有个,
	p6 := map[int]string{
		1: "test1",
	}
	fmt.Println("p6 :", p6)

	// 这个映射的键（key）是字符串类型，值（value）可以是任意类型，包括基本数据类型、结构体、数组、切片、映射等
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      "30",
		"hobby":    []string{"读书", "爬山"},
	}
	fmt.Println("map data :", res)

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)
}
