package main
import "fmt"

var a = 10
func main(){
	fmt.Println(a)
}

func init(){
	fmt.Println(a)
	a = 20
}
// The init function is a special function in Go that is automatically called before the main function. it is used to initialize variables or perform setup tasks before the program starts executing. In this case, the init function prints the value of "a" (which is 10) and then assigns a new value (20) to "a". When the main function is called, it prints the updated value of "a" (which is now 20). So, the output will be: