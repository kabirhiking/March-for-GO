//anonymous function
package main
import "fmt"

//strundard function
func add(x int, y int){
	fmt.Println("sum is", x+y)
}


func main(){
	//we invoke the function by passing the arguments
	add(2,3)


	//anonymous function
	//syntax: func (parameters) { function body }(arguments)
	//Imediately Invoked Function Expression (IIFE)
	func(a int, b int){
		c := a + b
		fmt.Println(c)
	}(4,5)
	
	a := 10 // expression

	//if expression
	if a > 0 { // if block
		fmt.Println("a is greater than 0")
	}
}

func init() {
		fmt.Println("i will call backe ASAP")
	}
