package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 = *&Person{}
	p1.Age = 10
	fmt.Println("P1=", p1)

	p2 := Person{Name: "kathy", Age: 24}
	fmt.Println("p2=", p2)

	var p3 = Person{Name: "kkkk", Age: 12}
	fmt.Println("p3=", p3)

	var p4 Person
	p4.Age = 30
	p4.Name = "kinwang"
	fmt.Println(p4)

	p5 := struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 100}
	fmt.Println("p5=", p5)
}
