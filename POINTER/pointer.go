package main

import "fmt"

//pass by value
//pass by reference

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {
	// pointer or address of memory (ram)

	// x := 20

	// p := &x // ampersand & => address of 

	// *p = 30

	// fmt.Println("x =", x)
// 	fmt.Println(p)  						 // p is the address of x
// 	fmt.Println("Value at the address", *p)  // * -> value at address

 arr := [3] int{1, 2, 3}
 print(&arr)
}


/*
2. Phases:
	1. compilation phase (compile time)
	2. execution phase (run time)


*********Compile Phase *********


### Code segment ###

print = func (numbers[3]int)
main = func() {...}



*/