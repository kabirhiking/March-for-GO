package main

import "fmt"

type User struct { // member variable or properties or fields
	// public variable
	name string
	age int
}



func main() {
	var user1 User

	user1 = User{ // instance or object or constructor
		name : "kabir",
		age : 30,
	}

	
	user2 := User{ // instance bananor prukriya ke bole instantiate
		name : "Abdullah",
		age : 25,
	}

	fmt.Println("Name:", user1.name)
	fmt.Println("Age", user1.age)
	fmt.Println("Name:", user2.name)
	fmt.Println("Age:", user2.age)


}