package main

import (
	"fmt"
	"example.com/mathlib" //importin custiom package
)
var a int = 20
var b int = 30


func main(){
	fmt.Println("Showing Custom Package Scope")
	mathlib.Add(a, b)
}