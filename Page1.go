package main

import (
	"fmt"
	"strings"
)

// In this Tutorial we are gointg to get farmiliar with the Go methods available
// for manipulating strings
// 1. Func compare: this function compares two strings and return
//an integer, the code is as follows:

// GO language program using String Compare function to compare 2 strings
func compare(a, b string) int{
	fmt.Println(strings.Compare("Nigeria", "niger"))
	fmt.Println(strings.Compare("niger", "Nigeria"))
}
func main(){
	re
}
//The strings.Compare returns 3 possible values
// which is 0(equality), 1(greater than) or -1(less than)