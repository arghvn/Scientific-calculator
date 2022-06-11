package main


// The below example is to write a Simple Calculator which supports arithmetic Operators such as Add, Subtract, Multiply, 
// Modulus, Division operations.
// First Read the input numbers from the user keyboard console, stored in variables number1 and number2.
// And also Asks the user to enter Operator from the console, stores it in the Operator symbol. 
// The user values are read from the console using the Scanln function.
// For example, If the user entered number1=20 and number2=10 and operator=+
// Matched Operator case executed using switch statements and returned the result.
// Finally, Print the output to the console using the Printf function.


import "fmt"

func main() {
    var operator string
    var number1, number2 int
    fmt.Print("Please enter First number: ")
    fmt.Scanln( & number1)
    fmt.Print("Please enter Second number: ")
    fmt.Scanln( & number2)
    fmt.Print("Please enter Operator (+,-,/,%,*):")
    fmt.Scanln( & operator)
    output: = 0
    switch operator {
        case "+":
            output = number1 + number2
            break
        case "-":
            output = number1 - number2
        case "*":
            output = number1 * number2
        case "/":
            output = number1 / number2
        case "%":
            output = number1 % number2
        default:
            fmt.Println("Invalid Operation")
    }
    fmt.Printf("%d %s %d = %d", number1, operator, number2, output)
}

// The output of the above program is



// output :

// Please enter First number: 20  
// Please enter Second number: 10  
// Please enter Operator (+,-,/,%,*):+  
// 20 %!c(string=+) 10 = 30  
  
// Please enter First number: 20  
// Please enter Second number: 10  
// Please enter Operator (+,-,/,%,*):+  
// 20 + 10 = 30  
  
// Please enter First number: 10  
// Please enter Second number: 5  
// Please enter Operator (+,-,/,%,*):/  
// 10 / 5 = 2  
  
// Please enter First number: 8  
// Please enter Second number: 2  
// Please enter Operator (+,-,/,%,*):%  
// 8 % 2 = 0  
  
// Please enter First number: 6  
// Please enter Second number: 2  
// Please enter Operator (+,-,/,%,*):*  
// 6 * 2 = 12  
  
// Please enter First number: 10  
// Please enter Second number: 2  
// Please enter Operator (+,-,/,%,*):-  
// 10 - 2 = 8 