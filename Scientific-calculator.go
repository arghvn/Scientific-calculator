package main

//This is a developing project

import (

 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
 "math"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 fmt.Print("Enter Number 1: ")
 number_1, _ := reader.ReadString('\n')

 fmt.Print("Enter Number 2: ")
 number_2, _ := reader.ReadString('\n')

 fmt.Print("Enter Operation: ")
 operation, _ := reader.ReadString('\n')

 number_1 = strings.Trim(number_1, "\n")
 number_2 = strings.Trim(number_2, "\n")

 num_1, _ := strconv.Atoi(number_1)
 num_2, _ := strconv.Atoi(number_2)

 switch strings.Trim(operation, "\n") {
 case "+":
  fmt.Println(Add(num_1, num_2))
 case "-":
  fmt.Println(Sub(num_1, num_2))
 case "*":
  fmt.Println(Mul(num_1, num_2))
 case "/":
  fmt.Println(Div(float32(num_1), float32(num_2)))
 }
var x float64
	x = math.Sin(x)
	fmt.Println(x)

var x float64
    x = math.cos(x)
	fmt.Println(x)

var x float64
    x = math.tan(x)
	fmt.Println(x)

var x float64
    x = math.cot(x)
	fmt.Println(x)


} 



	)
 // fmt.Println(number_1, number_2, operation)
}

func Add(a int, b int) int {
 return a + b
}

func Sub(a int, b int) int {
 return a - b
}

func Mul(a int, b int) int {
 return a * b
}

func Div(a float32, b float32) float32 {
 return a / b
}



func Factorial(number int) int {
	if (number == 1){
		return 1
	}
	return number* Factorial(number - 1)
}


//I have to learn the geen framework and convert the calculator 
//code to api with this framework.
//How to create an api ???
