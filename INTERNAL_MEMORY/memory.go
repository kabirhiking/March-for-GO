package main
import "fmt"



var a int = 10

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}


func main() {
	add(5, 4)// function call with arguments
	add(a, 3)// function call with global variable

}

func init() {
	fmt.Println("kabir")
}

//Internal memory -> code segment -> data segment -> heap segment -> stack segment
	// 1. code segment -> joto gual function ache shob code segment e jai.
	// 2. data segment -> joto global variable ache shob code segment e hai.
	// 3. heap segment -> protketa function jokhon execution hoi oy function er joto variable shob stack e jai.  dynamic memory allocation -> pointer, slice, map, channel, struct, interface
	// 4. stack segment -> local variable, function call, function parameter, function return value
