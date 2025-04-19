// package main
// import "fmt"


// func add(x int, y int){
//   sum := x + y
//   fmt.Println("sum:", sum)
// }


// func main() {
//   a :=5
//   b := 3

//   sum := a + b
//   fmt.Println("sum:", sum)

//   add(a, b)
//   add(6, 9)
// } 

package main
import "fmt"

func add(x int, y int) int{
  sum := x + y
  return sum
}

func getNumbers(num1 int, num2 int) (int, int){
  sum := num1 + num2

  mul := num1 * num2

  return sum, mul
}

func oiKire(){
  fmt.Println("modhu modhu")
}
func sayHi(name string){
  fmt.Println("welcome to go: ", name)
}

func main(){
  // a := 5
  // b := 7
 
  // sum := add(a, b)
  // fmt.Println("meyer sum:", sum)


  // p, q := getNumbers(a, b)
  
  // fmt.Println(p)
  // fmt.Println(q)

  oiKire()
  sayHi("kabir")
  

}