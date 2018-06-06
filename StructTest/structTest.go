package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	fmt.Println(Person{"Bob",20})
	fmt.Println(Person{name:"Alice",age:30})
	fmt.Println(Person{name:"Fred"})
	fmt.Println(&Person{name:"Ann",age:40})
	s:=Person{"Sean",50}
	fmt.Println(s.name)
	sp:=&s
	fmt.Println(sp.age)
	fmt.Println(*sp)
	//fmt.Println(sp.age)

	var p *int
	fmt.Printf("p:%v\n",p)
	var i int
	p=&i
	fmt.Printf("p:%v,%v\n",p,*p)
	*p=8
	fmt.Printf("p:%v,%v\n",p,*p)
	m:=[3]int{3,4,5}
	fmt.Printf("m:%v--%v,%v,%v\n",m,m[0],m[1],m[2])
	x:=[3]*int{&m[0],&m[1],&m[2]}
	fmt.Printf("x:%v,%v,%v\n",x[0],x[1],x[2])
	fmt.Printf("*x:%v,%v,%v\n",*x[0],*x[1],*x[2])
}
