package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	os.Setenv("FOO","1")
	fmt.Println("FOO:",os.Getenv("FOO"))
	fmt.Println("BAR:",os.Getenv("BAR"))
	fmt.Println()

	for _,e:=range os.Environ(){
		pair:=strings.Split(e,"=")
		println(pair[0],":",pair[1])
	}
}
