package main
import "fmt"

var a = 10

func main(){
	  age := 30

	  if age > 18 {
		a := 47
		fmt.Println(a)
	  }

	  fmt.Println(a)
}
//variable shadowing occurs when a variable in an inner scope has the same name as a variable in an outer scope. In this case, the inner variable "a" shadows the outer variable "a" within the if block. As a result, when we print "a" inside the if block, it refers to the inner variable, and when we print "a" outside the if block, it refers to the outer variable.