package introduction

import (
	"fmt"
)

func Variable() {
	 // explicitly declare the data type
    var number1 int = 10
    fmt.Println(number1)
    
    // assign a value without declaring the data type
    var number2 = 20
    fmt.Println(number2)
    
    // shorthand notation to define variable
    number3 := 30
    fmt.Println(number3)  

    // Changing value 
    // initial value
    number := 10
    fmt.Println("Initial number value", number) // prints 10
    
    // change variable value
    number = 100
    fmt.Println("The changed value", number)  // prints 100

    const lightSpeed = 299792458 // initial value
    // with constant we can't changed the value & can't be used with shorthand notation :=

    var salary1 float32
    var salary2 float64

    salary1 = 50000.5038829013333333

  // can store decimals with greater precision
    salary2 = 50000.50388290133333333333333

    fmt.Println(salary1) 
    fmt.Println(salary2)

    // Print statement 
    fmt.Print("this is print")
    fmt.Println("This is Println")
    currentAge := 21
    fmt.Printf("Age = %d", currentAge)
    // Integer %d
    // Float %g
    // string %s
    // bool %t

    var name = "John"
    age := 23
    
    fmt.Printf("%s is %d years old.", name, age)
    var name2 string

    // takes input value for name
    fmt.Print("Enter your name: ")
    fmt.Scan(&name2)
    
    fmt.Printf("Name: %s", name2)
    num1 := 6
    num2 := 2
    
    // + adds two variables
    sum := num1 + num2
    fmt.Printf("%d + %d = %d\n", num1, num2, sum)
    
    // - subtract two variables
    difference := num1 - num2
    fmt.Printf("%d - %d = %d\n",num1, num2,  difference)
    
    // * multiply two variables
    product := num1 * num2
    fmt.Printf("%d * %d is %d\n",num1, num2,  product)
}
