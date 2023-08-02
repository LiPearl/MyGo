package main

import "fmt"

func main() {
	//one dim
	var arr1 []int
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2}
	fmt.Println(arr4)

	//初始化特定位置数组 => [3 5 0 0 6]
	arr5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr5)

	//two dims
	var arr6 = [3][5]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {6, 7, 8, 9, 10}}
	fmt.Println(arr6)

	arr7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr7)

	//[[1 2 3] [4 5] []] 第二个[]可以为空，不可以为...
	arr8 := [3][]int{{1, 2, 3}, {4, 5}}
	fmt.Println(arr8)
}
