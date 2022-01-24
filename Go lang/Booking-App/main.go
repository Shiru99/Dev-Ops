package main

import (
	"Booking-App/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Global Variable vs Package Variable vs Local Variable
// Global variables - 1st letter capital
// var GlobalVariable string = "I'm Global Variable"

// Package level variables
var conferenceName = "'Go Conference'"

// conferenceName:=  "'Go Conference'"	// error for package level variable
const conferenceTickets = 50

var remainingTickets uint = 50

var waitGroup = sync.WaitGroup{}

func main() {

	for {

		/****** Variables & Constants ******/
		// var conferenceName = "'Go Conference'"
		// const conferenceTickets = 50

		// var remainingTickets = 50;
		// remainingTickets := 50 /* Only for var */
		// var remainingTickets uint = 50 /* Correct method */ // uint - unsigned int (+ve)

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
		fmt.Print("Enter your username : ")
		fmt.Scanln(&userName)
		fmt.Print("Enter your mail ID : ")
		fmt.Scanln(&mailId)
		fmt.Print("Enter number of tickets to book : ")
		fmt.Scanln(&userTickets)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Tickets remaining : %v\n", remainingTickets)
		// fmt.Printf("%v booked %v tickets\n", userName, userTickets)

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

		/****** Switch ******/
		// city := "Pune"

		// switch city {
		// case "Bangalore":
		// 	fmt.Println("welcome to Bangalore")
		// case "Mumbai", "Chennai":
		// 	fmt.Printf("welcome to %v\n", city)
		// case "Goa":
		// 	fmt.Println("welcome to Goa")
		// 	fmt.Println("You can visit nearby beach")
		// default:
		// 	fmt.Println("welcome to India")
		// }

		greetings("Hello, World!")
		fmt.Println(sum(88, 11))

		sum, sub, mul, div, rem := helper.Operations(11, 9)
		fmt.Printf("Summation : %v\nSubtraction : %v\nMultiplication : %v\nDivision : %v\nReminder : %v\n", sum, sub, mul, div, rem)

		/****** Maps ******/
		// bookedTickets := addUsers(userName, mailId, userTickets)

		// fmt.Printf("\n%v booked %v tickets\n", userName, userTickets)
		// for _, ticket := range bookedTickets {
		// 	fmt.Printf("\t%v %v, %v\n", ticket["firstName"], ticket["lastName"], ticket["age"])
		// }

		/****** Struct ******/
		bookedTickets_ := addUsersUsingStruct(userName, mailId, userTickets)

		fmt.Printf("\n%v booked %v tickets\n", userName, userTickets)
		fmt.Println(bookedTickets_)
		for _, ticket := range bookedTickets_ {
			fmt.Printf("\t%v %v, %v\n", ticket.firstName+" "+ticket.lastName, ticket.email, ticket.age)
		}

		/****** Concurrency & Synchronization ******/
		waitGroup.Add(1) // 1 - for 1 new thread
		go sendTickets(userName, mailId, userTickets, bookedTickets_)

		var toContinue bool
		fmt.Print("Want to book more tickets (true/false): ")
		fmt.Scanln(&toContinue)
		if !toContinue {
			fmt.Println("please wait! Shuting down background processes...")
			// this thread (main thread) will wait until all the threads in waitGroup are Done.
			waitGroup.Wait()
			fmt.Println("Thanks! for using our service. have a great day!")
			break
		} else {
			fmt.Println("Please book next ticket")
		}

	}
}

func greetings(greet string) {
	fmt.Println(greet, "from hellothere fun")
}

func addUsers(userName string, mailID string, userTickets uint) []map[string]string {

	var bookedTickets = make([]map[string]string, 0)

	fmt.Println("\nPlease add details for each ticket :")

	for i := 0; i < int(userTickets); i++ {

		/* NO mix datatypes for value/keys in map */
		// var ticketData map[string]string		// Declaration

		var ticketData = make(map[string]string) // Definition

		fmt.Println("\nUser", i+1, ": ")

		var firstName string = "default"
		var lastName string = "default"
		var age uint = 0

		fmt.Print("Enter firstname : ")
		fmt.Scanln(&firstName)
		fmt.Print("Enter lastname : ")
		fmt.Scanln(&lastName)
		fmt.Print("Enter age : ")
		fmt.Scanln(&age)

		fmt.Printf("seat confirmed for : %v %v, %d\n", firstName, lastName, age)
		ticketData["firstName"] = firstName
		ticketData["lastName"] = lastName
		ticketData["age"] = strconv.FormatUint(uint64(age), 10)

		bookedTickets = append(bookedTickets, ticketData)

	}

	return bookedTickets
}

type userDetails struct {
	firstName string
	lastName  string
	email     string
	age       uint
}

func addUsersUsingStruct(userName string, mailID string, userTickets uint) []userDetails {

	var bookedTickets = make([]userDetails, 0)

	fmt.Println("\nPlease add details for each ticket :")

	for i := 0; i < int(userTickets); i++ {

		fmt.Println("\nUser", i+1, ": ")

		var firstName string = "default"
		var lastName string = "default"
		var email string = "default@mail"
		var age uint = 0

		fmt.Print("Enter firstname : ")
		fmt.Scanln(&firstName)
		fmt.Print("Enter lastname : ")
		fmt.Scanln(&lastName)
		fmt.Print("Enter mail id : ")
		fmt.Scanln(&email)
		fmt.Print("Enter age : ")
		fmt.Scanln(&age)

		var ticketData = userDetails{
			firstName: firstName,
			lastName:  lastName,
			email:     email,
			age:       age,
		}
		bookedTickets = append(bookedTickets, ticketData)
		fmt.Printf("seat confirmed for : %v %v, %d\n", firstName, lastName, age)
	}
	return bookedTickets
}

func sendTickets(userName string, mailID string, userTickets uint, bookedtickets []userDetails) {
	var mailBody string = fmt.Sprintf("\nHi %v,\n\n"+"Your booked tickets are : \n\n", userName)
	// var mailBody string = "\nHi " + userName + ",\n\n" + "Your booked tickets are : \n\n"

	for _, ticket := range bookedtickets {
		mailBody = mailBody + "\t" + ticket.firstName + " " + ticket.lastName + ", " + ticket.email + ", " + strconv.FormatUint(uint64(ticket.age), 10) + "\n"
	}

	mailBody = mailBody + "\n\nThanks,\nTeam Hellothere"

	time.Sleep(20 * time.Second)
	fmt.Println("############################")
	fmt.Printf("Sending ticket details to email address : %v\n", mailID)
	fmt.Println(mailBody)
	fmt.Println("############################")
	fmt.Println("Done")
	// thread comes to end
	waitGroup.Done()
}
