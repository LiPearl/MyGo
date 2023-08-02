package main

import "fmt"

func main() {
	person := [3]string{"Tom", "Aaron", "john"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)
	for k, v := range person {
		fmt.Printf("person[%d]:%s\n", k, v)
	}
	for i := range person {
		fmt.Println(i)
	}
	for i := 0; i < len(person); i++ {
		if i == 2 {
			goto END
		}
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}
END:
	fmt.Println("end")
}
