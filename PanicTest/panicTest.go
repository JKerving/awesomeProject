package main

import (
	"fmt"
	"errors"
)

func testError() (err error){
	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("testError() 遇到错误:",r)
			switch x := r.(type) {
			case string:
				err=errors.New(x)
			case error:
				err=x
			default:
				err=errors.New("")
			}
		}
	}()
	panic("\"panic 错误\"")
	fmt.Println("抛出一个错误后继续执行代码")
	return nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("testError() 遇到错误:",r)
		var err error
		newA,ok:=r.(string)
		fmt.Println(newA,ok)
		switch x := r.(type) {
		case string:
			err=errors.New(x)
		case error:
			err=x
		default:
			err=errors.New("")
		}
		if err!=nil{
			fmt.Println("recover后的错误:",err)
		}
	}
}

func afterErrorfunc() {
	fmt.Println("遇到错误之后 func")
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName==nil{
		panic("runtime error:;first name cannot be nil")
	}
	if lastName==nil{
		panic("runtime error:last name cannot be nil")
	}
	fmt.Printf("%s %s\n",*firstName,*lastName)
	fmt.Println("returned normally from fullName")
}

func recoverName()  {
	if r := recover(); r != nil {
		fmt.Println("recovered from",r)
	}
}

func main() {
	//err:=testError()
	//if err!=nil{
	//	fmt.Println("main函数得到错误类型:",err)
	//}
	//afterErrorfunc()
	//defer fmt.Println("deferred call in main")
	//firstName:="Elon"
	//fullName(&firstName,nil)
	//fmt.Println("returned normally from main")
	i:=0
	defer fmt.Println(i)
	i++
	return
}

func divide(i, j int) {
	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("Recovered:",r)
		}
	}()
	fmt.Println(i/j)
}