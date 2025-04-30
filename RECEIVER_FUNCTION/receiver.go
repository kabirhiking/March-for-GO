package main

import "fmt"

type User struct { // member variable or properties or fields
	// public variable
	Name string
	Age int
}

func PrintUserDetails(usr User){
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// receiver function
func(usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// receiver function
func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	var user1 User

	user1 = User{ // instance or object or constructor
		Name : "kabir",
		Age : 30,
	}

	PrintUserDetails(user1)
	user1.printDetails()
	user1.call(7)
	

	
	// user2 := User{ // instance bananor prukriya ke bole instantiate
	// 	name : "Abdullah",
	// 	age : 25,
	// }

}



/*
2. Phases:
	1. compilation phase (compile time)
	2. execution phase (run time)


*********Compile Phase *********


### Code segment ###

User = type stuct {.....}
printUserDetails = func { usr User} {...}
PrintDetails = func() {....}// User
call = func() {...}
main = func() {...}



*/

