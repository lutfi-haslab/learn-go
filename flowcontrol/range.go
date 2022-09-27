package flowcontrol

import "fmt"

func Range(){
    // array of numbers
    numbers := [5]int{21, 24, 27, 30, 33}
    
    // using range to iterate over the elements of array
    for index, item := range numbers {
    fmt.Printf("numbers[%d] = %d \n", index, item)
    }

     // create a map
    subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
    fmt.Println("Marks obtained:")
    
    // using for range to iterate through the key-value pair
    for subject, marks := range subjectMarks {
    fmt.Println(subject, ":", marks)
    }
}