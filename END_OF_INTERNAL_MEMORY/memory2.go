package main
import "fmt"

const a = 10 // constant variable

var p = 100 // global variable

func call() {
add := func(x int, y int) {
	z := x + y
	fmt.Println(z)
}
add(5, 6) // function expression
add(p, a)
}

func main() {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("kabir")
}

/*
  phase
  1. complation phase -> syntax error, semantic error, logical error
  2. execution phase -> run time error, panic error, nil pointer dereference, index out of range, divide by zero, type assertion error, type conversion error, array out of bound error, slice out of bound error, map out of bound error


  go run main.go -> compile it -> main -> ./main
  go build main.go -> compile it -> main
  ./main -> run it


## code segment ekta binary file er modhe thake jokhon ./main diye file run kori tokhon ram er modhe ashe code segment er por first to last read kore. erpor global variable gual Data segment e ashe. joto function ache shob stack frame occupied kore run hoy(stack frame e jegula thake egula local scope. Code segment er jinish potro read only change kora jayna but data segment er jinish pati chang kora jay).

*/


/*
   #### Binary file ###
   	010101010110
	010101010110
	010101010110
	010101010110
	010101010110
	010101010110
*/	

