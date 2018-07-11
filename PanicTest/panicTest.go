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

func main() {
	err:=testError()
	if err!=nil{
		fmt.Println("main函数得到错误类型:",err)
	}
	afterErrorfunc()
}

func divide(i, j int) {
	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("Recovered:",r)
		}
	}()
	fmt.Println(i/j)
}