package main

import (
	"sort"
	"fmt"
)

func main() {
	strs:=[]string{"c","b","a"}
	sort.Strings(strs)
	fmt.Println(strs)
	ints:=[]int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints:",ints)

	s:=sort.IntsAreSorted(ints)
	fmt.Println(s)
}
