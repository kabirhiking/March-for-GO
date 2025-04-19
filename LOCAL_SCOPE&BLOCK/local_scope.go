package main
import "fmt"

var a int = 20
var b int = 30

// block -> { }

func main(){
	x := 18

	if x >= 18 {
		p := 30
		fmt.Println("I am mature boy")
		fmt.Println("I have", p, "girlfriends")
	}

	fmt.Println("I have", a, "girlfriends")
}

//local scope e 3 types er scope ache -> 1.if scope 2.function scope 3.switch scope
