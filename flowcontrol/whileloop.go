package flowcontrol

import "fmt"

func WhileLoop(){
    // Program to create a multiplication table of 5 using while loop
    multiplier := 1

    for multiplier <= 10 {
        product := 5 * multiplier
        fmt.Printf("5 * %d = %d\n", multiplier, product)
        multiplier++
    }
}