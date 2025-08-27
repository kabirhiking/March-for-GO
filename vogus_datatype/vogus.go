package main
import "fmt"




func main() {
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	var e float32 = 3.4e+38
	var f float64 = 1.7e+308
	var boolVar bool = true
	var str string = "Hello, World!"
	var runeVar rune = 'A' // Unicode code point for 'A' love or other sign 



	fmt.Printf("%d", a)
	fmt.Printf("%d", b)
	fmt.Printf("%d", c)
	fmt.Printf("%d", d)
	fmt.Printf("%f", e)
	fmt.Printf("%f", f)
	fmt.Printf("%t", boolVar)
	fmt.Printf("%s", str)
	fmt.Printf("%c", runeVar)

}