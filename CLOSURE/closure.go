package main
import "fmt"

const a = 10 // constant variable

var p = 100 // global variable

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
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
  1. complation phase(compile time) -> syntax error, semantic error, logical error
  2. execution phase(run time) -> run time error, panic error, nil pointer dereference, index out of range, divide by zero, type assertion error, type conversion error, array out of bound error, slice out of bound error, map out of bound error


  go run main.go -> compile it -> main -> ./main
  go build main.go -> compile it -> main
  ./main -> run it


## code segment ekta binary file er modhe thake jokhon ./main diye file run kori tokhon ram er modhe ashe code segment er por first to last read kore. erpor global variable gual Data segment e ashe. joto function ache shob stack frame occupied kore run hoy(stack frame e jegula thake egula local scope. Code segment er jinish potro read only change kora jayna but data segment er jinish pati chang kora jay).

Closure -> ekta function er modhe onno ekta function thake. jokhon outer function ta call hoy tokhon inner function ta call hoy. inner function ta outer function er variable gual access korte pare. closure er maddhome amra data hiding korte pari. closure er maddhome amra encapsulation o korte pari.
escap e analysis -> jokhon amra function ta call kori tokhon stack frame e function ta thake. jokhon amra function ta return kori tokhon stack frame ta remove hoye jay. kintu jokhon amra inner function ta outer function er variable gual access kori tokhon stack frame ta remove hoye na. karon inner function ta outer function er variable gual access korte pare. tai closure er maddhome amra data hiding korte pari.
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

