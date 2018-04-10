package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark:=Student{Human{"Mark",18,80},"Computer Science"}
	//访问相应的字段
	fmt.Println("His name is ",mark.name)
	fmt.Println("His age is ",mark.age)
	fmt.Println("His weight is ",mark.weight)
	fmt.Println("His speciality is ",mark.speciality)
}