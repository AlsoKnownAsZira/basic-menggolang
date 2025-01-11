package main 
import "fmt" 

func main() {
fmt.Println("int data type")
var intNum int = 10
fmt.Println(intNum)

fmt.Println("Floating numbers, must be declared as float32 or float64")
var floatNum float32 = 22.3
fmt.Print(floatNum)

fmt.Println("Aritmetical operations")
var a int = 10
var b int = 20
var sum int = a + b
fmt.Println(sum)

fmt.Println("Casting data types (Go does no support cross data type operations)")
var x float32 = 99.1 
var cast = a + int(x)
fmt.Println(cast)

fmt.Println("String data type")
var str string = "Hello\nWorld"
var str2 string = "Hello" + " "+"World"
var str3 string = `Hello
World`
fmt.Println(str)
fmt.Println(str2)
fmt.Println(str3)

fmt.Println("Declaration types")
var name string = "John"
var name2 = "Doe"
name3 := "Jane"

var age1,age2 int = 20, 30
var age3, age4 = 40, 50

fmt.Println(name + " " + name2 + " " + name3)
fmt.Println(age1, age2, age3, age4)

fmt.Println("Constants")
const pi float32 = 3.14159
const pi2 = 3.14159
fmt.Println(pi, pi2)
}