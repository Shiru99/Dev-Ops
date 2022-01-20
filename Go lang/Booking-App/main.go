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
	fmt.Println("Reserve your seat quickly, only", remainingTickets, "out of", conferenceTickets, "tickets are remaining")

	/****** Data-type ******/
	fmt.Printf("conferenceName - %T, conferenceTickets - %T, remainingTickets - %T\n", conferenceName, conferenceTickets, remainingTickets)

	/****** Pointers ******/
	fmt.Printf("variable - %v, variable type - %T \n", conferenceName, conferenceName)
	fmt.Printf("variable address - %v, variable address type - %T \n", &conferenceName, &conferenceName)

	var userName string
	var userTickets uint

	/******  Input statements ******/
	fmt.Print("Enter your username : ")
	fmt.Scanln(&userName)
	fmt.Print("Enter number of tickets to book : ")
	fmt.Scanln(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v booked %v tickets\n", userName, userTickets)
	fmt.Printf("Tickets remaining : %v\n", remainingTickets)

}
