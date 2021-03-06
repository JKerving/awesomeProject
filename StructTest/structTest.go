package main

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct {
	name string
	id int
	score int
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

	var n [3]*int
	n = x
	fmt.Printf("n:%v,%v,%v\n",n[0],n[1],n[2])
	y:=[]*[3]*int{&x}
	fmt.Printf("y:%v,%v\n",y,y[0])
	fmt.Printf("*y[0]:%v\n",*y[0])
	fmt.Printf("*y[][]:%v,%v,%v\n",*y[0][0],*y[0][1],*y[0][2])

	var st Student
	st.name="Chinese"
	st.id=666333
	st.score=100
	fmt.Println(st)
	var ptr *[]Student
	fmt.Println(ptr)
	ptr=new([]Student)
	fmt.Println(ptr)
	*ptr=make([]Student,3,100)
	fmt.Println(ptr)
	stu:=[]Student{{"China",3333,66},{"Chinese",4444,77},{"Chince",5555,88}}
	fmt.Println(stu)
}
