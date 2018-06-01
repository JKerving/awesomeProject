package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}
}

func zero(z *int) {
	fmt.Println(z)
	*z = 10
}

func main() {
	var a [5]int
	fmt.Println("emp:",a)
	 a[4]=100
	 fmt.Println("set:",a)
	 fmt.Println("get:",a[4])

	 fmt.Println("len:",len(a))

	 b:=[5]int{1,2,3,4,5}
	 fmt.Println("dc1:",b)

	 var twoD [2][3]int

	 for i:=0;i<2;i++{
	 	for j:=0;j<3;j++{
	 		twoD[i][j]=i+j
		}
	 }

	 fmt.Println("2d:",twoD)
}
