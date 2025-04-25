// package main
// import "fmt"

// func add(x int, y int){ // fucntion parameters 
// 	c := x + y
// 	fmt.Println(c)
// }

// func main() {
	
	
// 	add(5, 5) //functin call with arguments
// }

//Higher order function
// package main
// import "fmt"

// func processOperation(a int, b int, op func(p int, q int)) func(p int, q int) {
// 	op(a,b)
// 	return add
// }

// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }

// func main() {
// 	processOperation(2, 5, add)
// }

package main
import "fmt"

func call() func(int, int) {
	return add 
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	sum := call()
	sum(4, 3) // function expression
}



/*
	1. parameter vs argument
	2. First order function -> Normal jinish(variable) niye kaj kore
		i. standard function or named function
		ii. anonymous function 
		iii. IIFE (Immediately Invoked Function Expression)
		iv. function expression
	3. Higher order function or (first class function-> ) -> je function kono function ke receive kore as a parameter hishebe or return kore as a value or both
	4. Callback function -> function ke parameter hisebe pass kora hoyeche
	5. First class citizen -> function ke first class citizen bola hoyeche karon function ke variable hisebe assign kora jabe, function ke parameter hisebe pass kora jabe, function ke return kora jabe, function ke store kora jabe data structure e, function ke call kora jabe, function ke execute kora jabe, function ke invoke kora jabe, function ke use kora jabe, function ke create kora jabe, function ke define kora jabe, function ke declare kora jabe, function ke implement kora jabe, function ke use kora jabe, function ke call kora jabe, function ke execute kora jabe, function ke invoke kora jabe, function ke use kora jabe, function ke create kora jabe, function ke define kora jabe, function ke declare kora jabe, function ke implement kora jabe.
	
	functional paradigm -> haskell, racket, lisp, clojure, scala, ocaml, erlang, f#, elixir, etc.

	1. first class order logic -> first class function -> first class citizen -> first class object -> first class data type -> first class data structure
	2. higher order logic -> higher order function -> higher order citizen -> higher order object -> higher order data type -> higher order data structure

	#### logic ####

	1. object (people, animal, thing, etc.) -> object oriented programming (oop) -> object oriented logic (ool) -> object oriented paradigm (oop) -> object oriented design (ood) -> object oriented analysis (ooa) -> object oriented methodology (oom) -> object oriented approach (ooa) -> object oriented software engineering (oose) -> object oriented software development (oosd) -> object oriented software architecture (oosa) -> object oriented software design (oods) -> object oriented software analysis (ooda) -> object oriented software methodology (ooma) -> object oriented software approach (ooa) -> object oriented software engineering process (oosep) -> object oriented software development process (oosdp) -> object oriented software architecture process (oosap) -> object oriented software design process (oodspp) -> object oriented software analysis process (oodap) -> object oriented software methodology process (oomap) -> object oriented software approach process (ooap)
	2. property(color, student id, name, etc.) -> property oriented programming (pop) -> property oriented logic (pol) -> property oriented paradigm (pop) -> property oriented design (pod) -> property oriented analysis (poa) -> property oriented methodology (pom) -> property oriented approach (poa) -> property oriented software engineering (pose) -> property oriented software development (posd) -> property oriented software architecture (posa) -> property oriented software design (posd) -> property oriented software analysis (posa) -> property oriented software methodology (posm) -> property oriented software approach (posa) -> property oriented software engineering process (posep) -> property oriented software development process (posdp) -> property oriented software architecture process (posap) -> property oriented software design process (posdp) -> property oriented software analysis process (posaap) -> property oriented software methodology process (posmp) -> property oriented software approach process (poap)
	3. Relationship (student, teacher, etc.) -> relationship oriented programming (rop) -> relationship oriented logic (rol) -> relationship oriented paradigm (rop) -> relationship oriented design (rod) -> relationship oriented analysis (roa) -> relationship oriented methodology (rom) -> relationship oriented approach (roa) -> relationship oriented software engineering (rose) -> relationship oriented software development (rosd) -> relationship oriented software architecture (rosa) -> relationship oriented software design (rosd) -> relationship oriented software analysis (rosa) -> relationship oriented software methodology (rosm) -> relationship oriented software approach (rosa) -> relationship oriented software engineering process (rosep) -> relationship oriented software development process (rosdp) -> relationship oriented software architecture process (rosap) -> relationship oriented software design process (rosdp) -> relationship oriented software analysis process (rosaap) -> relationship oriented software methodology process (rosmp) -> relationship oriented software approach process (roap)

	Tutul is a student 
	Apple is red
	Tutul is taller than Rakib

	### First order logic ###

	Rule: All customers must pay their pizza bills.
	      All students must wear their uniforms.

	### Higher order logic ### 

	Any rule that appllies to all customers must also apply to tutul.const

	a rule : all customers must pay tips to the waiters.


*/

