package flowcontrol

import "fmt"

func IfElse(){
    number1 := 12
    number2 := 20
    
    if number1 == number2 {
    fmt.Printf("Result: %d == %d", number1, number2)
    } else if number1 > number2 {
    fmt.Printf("Result: %d > %d", number1, number2)
    } else {
    fmt.Printf("Result: %d < %d", number1, number2)     
    }
    fmt.Println()
}