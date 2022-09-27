// Program to print the day of the week using  switch case

package flowcontrol
import ("fmt")

func Switch() {
  dayOfWeek := 3

  switch dayOfWeek {

    case 1:
      fmt.Println("Sunday")
   	 
    case 2:
      fmt.Println("Monday")
   	 
    case 3:
      fmt.Println("Tuesday")
   	 
    case 4:
      fmt.Println("Wednesday")
   	 
    case 5:
      fmt.Println("Thursday")
   	 
    case 6:
      fmt.Println("Friday")
   	 
    case 7:
      fmt.Println("Saturday")
   	 
    default:
      fmt.Println("Invalid day")
  }

    dayOfWeekString := "Sunday"

    switch dayOfWeekString {
    case "Saturday", "Sunday":
      fmt.Println("Weekend")
    
    case "Monday","Tuesday","Wednesday","Thursday","Friday":
      fmt.Println("Weekday")
    
    default:
      fmt.Println("Invalid day")
    }
}