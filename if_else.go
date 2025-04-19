// package main

// import "fmt"

// func main(){
// 	age := 18

// 	if age > 18 {
// 		fmt.Println("Your are eligible to be married")
// 	} else if age <= 18{
// 		fmt.Println("You are not eligible to be married")

// 	} else if age == 18{
// 		fmt.Println("Yor are just a teenager, not eligible to be married")

// 	}
// 	// >, <, =>, <=, ==,
// 	// and => &&
// 	// or => ||
// 	// not => !
// 	// != not equal
// }

package main

import "fmt"

func main() {
	age := 20
	sex := "male"

	if age == 20 && sex == "male" {
		fmt.Println("You are ready to be married")
	}
}

// >, <, =>, <=, ==,
// and => &&
// or => ||
// not => !
// != not equal
