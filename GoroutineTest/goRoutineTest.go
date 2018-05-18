package main

import "fmt"

var chanCap int = 5
var ch1 chan int = make(chan int,1)
var ch2 chan int = make(chan int,1)
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

type Rect struct {
	width,length float64
}

type Phone struct {
	price int
	color string
}

func (phone Phone) ringing()  {
	println("Phone is ringing...")
}

type IPhone struct {
	Phone
	model string
}

func (rect Rect) area() float64 {
	return rect.width*rect.length
}

func doubleArea(rect Rect) float64 {
	rect.width *=2
	rect.length *=2
	return rect.width*rect.length
}

func main() {
	//select {
	//case getChan(0)<-getNumber(2):
	//	fmt.Println("1th case is selected")
	//case getChan(1)<-getNumber(3):
	//	fmt.Println("2th case is selected")
	//default:
	//	fmt.Println("default")
	//}
	////caseTest()
	//c:=make(chan int)
	//quit:=make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit<-0
	//}()
	//fibonacci(c,quit)
	var rect = Rect{100,200}
	println(doubleArea(rect))
	println("width:",rect.width,"length:",rect.length)

	var p IPhone
	p.price=5888
	p.color="Black"
	p.model = "iPhone 5"
	p.ringing()
	println("I have a iPhone")
	println("Price:",p.price)
	println("Color",p.color)
	println("Model",p.model)
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n",i)
	return chs[i]
}

func caseTest()  {
	ch := make(chan int,chanCap)
	for i:=0;i<chanCap;i++{
		select {
		case ch<-1:
		case ch<-2:
		case ch<-3:
		}
	}

	for i:=0;i<chanCap;i++{
		fmt.Printf("%v\n",<-ch)
	}
	close(ch)
}

func fibonacci(c, quit chan int) {
	x,y:=1,1
	for{
		select {
		case c <- x:
			x,y=y,x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
