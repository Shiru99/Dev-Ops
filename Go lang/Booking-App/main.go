package main

import "fmt"

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

	programmingLang = append(programmingLang, "Go", "Java", "Dart", "Python", "Javascript")
	fmt.Println(programmingLang)
	fmt.Println(programmingLang[1])
	// fmt.Println(programmingLang[9])	/* panic: runtime error: index out of range [9] with length 5 */
	fmt.Println(len(programmingLang))

}
