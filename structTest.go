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
	//修改对应备注信息
	mark.speciality="AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ",mark.speciality)
	//修改年龄信息
	fmt.Println("Mark become old")
	mark.age=22
	fmt.Println("His age is ",mark.age)
	//修改体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight+=60
	fmt.Println("His weight is ",mark.weight)
}