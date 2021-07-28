package main

import (
	"fmt"
	"strings"
)

// In this Tutorial we are gointg to get farmiliar with the Go methods available
// for manipulating strings
// 1. Func compare: this function compares two strings and return three values
//an integer, the code is as follows:

// GO language program using String Compare function to compare 2 strings

func main(){	
		fmt.Println(strings.Compare("Nigeria", "niger"))
		fmt.Println(strings.Compare("niger", "Nigeria"))	
}
//The strings.Compare returns 3 possible values
// which is 0(equality), 1(greater than) or -1(less than)

//Func Contains takes a string and a substring as argument
 func Contains(s, substr) bool

 //Func ContainsAny returns a bool telling if any part of the chars is contained
 // In the string, unlike Contains that checks all character in the substring
 func ContainsAny(s, chars string) bool