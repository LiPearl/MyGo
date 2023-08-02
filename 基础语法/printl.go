package main

import "fmt"

// print输出不换行
//println 输出换行
//%v以原有值的形式输出

func main() {
	const name string = "tom"
	const age, hight = 20, 170.1
	fmt.Print("sdsjiosdf")
	fmt.Println("--------")
	fmt.Printf("name=%s,age=%d, height=%v\n", name, age, fmt.Sprintf("%.2f", 180.567))
}
