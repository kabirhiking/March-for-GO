package main
import "fmt"

var a int = 20
var b int = 30

func add(x int, y int){
	z := x +y
	fmt.Println(z)
}

func main(){
	p := 30
	q := 40

	add(p,q)
	
	add(a, b)

	add(a, p)   // z is not defined in this scope // z variableta ami jodi use korte na pari that means scope nai otherwise scope ache.

}
//Gloabal variable scope: global variable scope is the entire file.