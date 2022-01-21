package main

import (
	"fmt"
	"strings"
)

func main() {

	/****** Variables & Constants ******/
	var conferenceName = "'Go Conference'"
	const conferenceTickets = 50

	// var remainingTickets = 50;
	// remainingTickets := 50 /* Only for var */
	var remainingTickets uint = 50 /* Correct method */ // uint - unsigned int (+ve)

	/****** Print statements ******/
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println("Reserve your seat quickly, only", remainingTickets, "out of", conferenceTickets, "tickets are remaining")

	/****** Data-type ******/
	// fmt.Printf("conferenceName - %T, conferenceTickets - %T, remainingTickets - %T\n", conferenceName, conferenceTickets, remainingTickets)

	/****** Pointers ******/
	// fmt.Printf("variable - %v, variable type - %T \n", conferenceName, conferenceName)
	// fmt.Printf("variable address - %v, variable address type - %T \n", &conferenceName, &conferenceName)

	var userName string = "default"
	var mailId string = "default@de"
	var userTickets uint = 0

	/******  Input statements ******/
	// fmt.Print("Enter your username : ")
	// fmt.Scanln(&userName)
	// fmt.Print("Enter your mail ID : ")
	// fmt.Scanln(&mailId)
	// fmt.Print("Enter number of tickets to book : ")
	// fmt.Scanln(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v booked %v tickets\n", userName, userTickets)
	fmt.Printf("Tickets remaining : %v\n", remainingTickets)

	/****** Arrays ******/
	// var temp = [50]string{"t1", "t2", "t3"}
	var bookings [50]string
	bookings[0] = "*"
	bookings[25] = userName + "." + mailId
	bookings[49] = "*"

	// fmt.Println(bookings[25])
	// fmt.Println(bookings)
	// fmt.Printf("Array length - %T\n",bookings)
	// fmt.Println(len(bookings))

	/****** Dynamic Arrays - Slice ******/
	var programmingLang []string

	programmingLang = append(programmingLang, "Go lang", "Java lang", "Dart lang", "Python lang", "Javascript lang")
	// fmt.Println(programmingLang)
	// fmt.Println(programmingLang[1])
	// // fmt.Println(programmingLang[9])	/* panic: runtime error: index out of range [9] with length 5 */
	// fmt.Println(len(programmingLang))

	/****** Loops ******/
	// var num uint = 0
	// for i := 0; i < 5; i++ {
	// 	fmt.Print("Enter your number : ")
	// 	fmt.Scanln(&num)
	// 	fmt.Printf("Your input number was : %v\n", num)
	// }

	// for index, lang := range programmingLang{
	// 	var langName = strings.Fields(lang)[0]
	// 	fmt.Printf("%v - %v\n", index, langName)
	// }

	// _ - blank identifier - used to ignore the value of a variable
	for _, lang := range programmingLang {
		var langName = strings.Fields(lang)[0]
		fmt.Printf("%v\n", langName)
	}

	var boolTemp bool = true
	fmt.Printf("value of boolTemp : %v\n", boolTemp)

	// infinite loop
	// for{ } == for true { }		// will run forever

	// var num1 uint = 0
	// // for with condition, as long as condition is true for loop works
	// for boolTemp {
	// 	fmt.Print("Enter your number : ")
	// 	fmt.Scanln(&num1)

	// 	if num1 > 99 {
	// 		fmt.Printf("Sorry, Can't handle 3 digit number\n")
	// 		break
	// 	} else if num1 > 9 {
	// 		fmt.Printf("Please enter single digit number - %v\n", num1)
	// 		continue
	// 	} else {
	// 		if num1==9{
	// 			boolTemp = false
	// 		}
	// 		fmt.Printf("Your input number was : %v\n", num1)
	// 	}
	// }

	city := "Pune"

	switch city {
	case "Bangalore":
		fmt.Println("welcome to Bangalore")
	case "Mumbai","Chennai":
		fmt.Printf("welcome to %v\n",city)
	case "Goa":
		fmt.Println("welcome to Goa")
		fmt.Println("You can visit nearby beach")
	default:
		fmt.Println("welcome to India")
	}

}
