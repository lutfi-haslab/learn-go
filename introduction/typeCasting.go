package introduction

import "fmt"

// Adalah cara mengubaha tipe data variable
func TypeCasting(){
   // variable of float type
    var floatValue float32 = 9.8
    
    // convert float to int
    var intValue int = int(floatValue)
    fmt.Println(intValue)
}