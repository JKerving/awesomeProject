package main

import "fmt"

var chanCap int = 5
var ch1 chan int = make(chan int,1)
var ch2 chan int = make(chan int,1)
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChan(0)<-getNumber(2):
		fmt.Println("1th case is selected")
	case getChan(1)<-getNumber(3):
		fmt.Println("2th case is selected")
	default:
		fmt.Println("default")
	}
	caseTest()
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
