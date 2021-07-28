package main

import(
	"fmt"
)
//Printf does not help with \n unlike Println
//%v for variable
//%q for quotes
//%T prints the type of varianble
//%f prints float type variables

func main(){
	var name string = "Nnamdi"
	state := 24
	fmt.Printf("My Name is %v, I hail from, %T", name, state)

}