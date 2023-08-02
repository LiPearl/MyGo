package main

import "fmt"

// 07-31记录
// 如果不用new进行初始化，会报错 => "runtime error: invalid memory address or nil pointer dereference"
// 指针作为引用类型需要初始化后才会拥有内存空间，才可以进行赋值
//https://blog.csdn.net/weixin_44211968/article/details/121742973?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-121742973-blog-126948669.235^v38^pc_relevant_sort_base1&spm=1001.2101.3001.4242.2&utm_relevant_index=4

func main() {

	//声明指针变量a,并未进行初始化
	var a *int

	//用new，对a进行初始化
	//new函数得到一个类型的指针，并且该指针的值为该类型的零值
	a = new(int)
	*a = 100
	fmt.Println(*a)

	//make函数只用于slice、map、chan的内存创建
	//返回的还是这三个引用类型本身
	var mymap = make(map[int]string, 5)
	mymap[0] = "mylove"
	mymap[1] = "shelove"
	fmt.Println(mymap)

	var sli_0 = []int{} //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_0), cap(sli_0), sli_0)

	var sli_1 = make([]int, 2, 9)
	fmt.Printf("len = %d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	var sli_2 = make([]int, 5, 10)
	fmt.Printf("len = %d cap=%d slice=%v\n", len(sli_2), cap(sli_2), sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)
	fmt.Println("sli[0:3] ==", sli_4[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4[0:3]), cap(sli_4[0:3]), sli_4[0:3])

	//[A:B:C] cap: C-A是新切片的容量 c<=len(arr)  len:B-A
	fmt.Println("sli[0:3:4] ==", sli_4[1:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4[1:3:4]), cap(sli_4[1:3:6]), sli_4[1:3:4])
}

/*
100
map[0:mylove 1:shelove]
len=0 cap=0 slice=[]
len = 2 cap=9 slice=[0 0]
len = 5 cap=10 slice=[0 0 0 0 0]
len=6 cap=6 slice=[1 2 3 4 5 6]
len=5 cap=5 slice=[1 2 3 4 5]
sli_4[1] 2
sli_4[0:2] [1 2]
sli[0:3] == [1 2 3]
len=3 cap=5 slice=[1 2 3]
sli[0:3:4] == [1 2 3]
len=3 cap=4 slice=[1 2 3]
*/
