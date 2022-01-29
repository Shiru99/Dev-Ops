package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	age         int
	contactInfo contactInfo
}

type contactInfo struct {
	email string
	phone int
}

func main() {
	john := person{"John", "Doe", 20, contactInfo{"John.Doe@gmail.com", 1234567890}}
	fmt.Println(john)

	jenny := person{firstName: "Jenny", lastName: "Doe", contactInfo: contactInfo{email: "Jenny.Doe@gmail.com"}}
	fmt.Println(jenny)

	var james person
	james.firstName = "James"
	james.contactInfo.phone = 9876543210

	// fmt.Printf("%v\n", james)
	// fmt.Printf("%+v\n", james)
	james.print()

	/************** Pointers(*) & Addresses(&) **************/

	/*
		pointerOfJames := &james
		pointerOfJames.updateName("Jimmy")

			ShortCut : james.updateName("Jimmy")
	*/
	james.updateName("Jimmy")
	james.print()

	name := "James Bond"
	fmt.Println(&name)
	fmt.Println(*&name)

	/************** Pass By Value **************/

	fmt.Println(name) // "James Bond"
	funWithName(name) // pass by value
	fmt.Println(name) // "James Bond"

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums) // [1 2 3 4 5]
	funWithNums(nums)
	fmt.Println(nums) // [0 2 3 4 5]

	/*
		Everything in Go is passed by value, slices too. But a slice value is a header, describing a contiguous section of a backing array, and a slice value only contains a pointer to the array where the elements are actually stored. The slice value does not include its elements (unlike arrays).

		value types : int float string bool structs (Use pointers to change this thing in a function)
		reference types : slices maps channels pointers functions
	*/

		/************ Maps **************/
		colors := map[string]string{
			"Red":   "#ff0000",
			"Green": "#4bf745",
			"Blue":  "#0000ff",
			"White": "#ffffff",
			"Black": "#000000",
		}
	
		// To delete a key from map
		delete(colors, "Yellow")
		delete(colors, "Black")
	
	
		for key, value := range colors {
			fmt.Printf("%v : %v\n", key, value)
		}
}

// Receiver function
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Receiver function
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func funWithName(newName string) {
	newName = "blah blah"
}

func funWithNums(numbers []int) {
	numbers[0] = 0
	numbers[1] = 0
	numbers[2] = 0
	numbers[3] = 0
	numbers[4] = 0
}