package main

import "fmt"

/*如果一个类里面有一个字段叫做phone，而另外一个类也有一个字段叫做phone，这种情况下的访问方式是什么*/
type Human struct {
	name string
	age int
	phone string	//Human类型拥有的字段
}

type Employee struct {
	Human	//匿名字段
	speciality string
	phone      string	//雇员的phone字段
}

func main() {
	Bob:=Employee{Human{"Bob",34,"18910599692"},"Designer","333-222"}
	fmt.Println("Bob's work phone is:",Bob.phone)
	//如果需要访问Human的phone字段
	fmt.Println("Bob's personal phone is ",Bob.Human.phone)
}
