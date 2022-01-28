package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactDetails contactInfo
}

type contactInfo struct {
	email   string
	phone  int
}

func main()  {
	john := person{"John", "Doe", 20, contactInfo{"John.Doe@gmail.com", 1234567890}}
	fmt.Println(john)

	jenny := person{firstName: "Jenny", lastName: "Doe", contactDetails: contactInfo{email: "Jenny.Doe@gmail.com"}}
	fmt.Println(jenny)

	var james person
	james.firstName = "James"
	james.contactDetails.phone = 9876543210

	fmt.Printf("%v\n",james)
	fmt.Printf("%+v\n", james)
}