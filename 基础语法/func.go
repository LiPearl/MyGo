package main

import "fmt"

func main() {
	var arr1 = [6]int{1, 2, 3}
	modifyArr1(&arr1)
	modifyArr2(arr1)
	fmt.Println(arr1)
}

// 指针传递
func modifyArr1(arr1 *[6]int) {
	arr1[4] = 10
	arr1[0] = 20
}

// 值传递
func modifyArr2(arr2 [6]int) [6]int {
	arr2[1] = 111
	arr2[2] = 222
	return arr2
}
