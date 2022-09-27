package flowcontrol

import "fmt"

func Boolean(){
    var boolTrue bool = true
    var boolFalse bool = false
    
    println("The boolean values are", boolTrue, "and", boolFalse)

    number1 := 6
    number2 := 12
    number3 := 6
    var result bool
    
    // equal to operator
    result = (number1 == number2)
    fmt.Printf("%d == %d returns %t \n", number1, number2, result)
    
    // not equal to operator
    result = (number1 != number2)
    fmt.Printf("%d != %d returns %t \n", number1, number2, result)
    
    // less than operator
    result = (number1 > number2)
    fmt.Printf("%d > %d returns %t \n", number1, number2, result)
    
    // greater than operator
    result = (number1 < number2)
    fmt.Printf("%d < %d returns %t \n", number1, number2, number1 < number2)

    // returns false because number1 > number2 is false
    result = (number1 > number2) && (number1 == number3);
    fmt.Printf("Result of AND operator is %t \n", result);
    
    // returns true because number1 == number3 is true
    result = (number1 > number2) || (number1 == number3);
    fmt.Printf("Result of OR operator is %t \n", result);
    
    // returns false because number1 == number3 is true
    result = !(number1 == number3);
    fmt.Printf("Result of NOT operator is %t \n", result);
}