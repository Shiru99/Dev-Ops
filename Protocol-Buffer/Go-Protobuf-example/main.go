package main

import (
	"fmt"
	pb "github.com/Shiru99/Go-Protobuf-example/src/mypb"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// hello()
	complex()
}

func hello() {
	person := *getPerson()

	fmt.Println(person)
	fmt.Println(person.GetName())

	// JSON
	jsonString := protojson.Format(&person)
	fmt.Println(jsonString)

	// Marshal to JSON
	out, err := protojson.Marshal(&person) // Marshal returns a byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// Unmarshal from JSON
	var person2 pb.Person
	b := []byte(jsonString)
	err = protojson.Unmarshal(b, &person2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person2.GetName())
	person2.Name = "Dr. Joe Doe"
	fmt.Println(person2.GetName())
}

func complex() {
	appointmentBook := *getAppointmentBook()
	// fmt.Println(appointmentBook)
	fmt.Println(appointmentBook.GetPatients())
	fmt.Println(appointmentBook.GetContactDetails())
	fmt.Println(appointmentBook.GetAppointmentStatus())
	fmt.Println(appointmentBook.GetReason())
	fmt.Println(appointmentBook.GetIsCancelled())

	/* Marshal to JSON */
	out, err := protojson.Marshal(&appointmentBook) // Marshal returns a byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	/* multi-line */
	option := protojson.MarshalOptions{Multiline: true}
	out, err = option.Marshal(&appointmentBook) // Marshal returns a byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	/* Unmarshal from JSON */
	var appointmentBook2 pb.AppointmentBook
	b := []byte(string(out))
	err = protojson.Unmarshal(b, &appointmentBook2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(appointmentBook2)

	/* Discard unknown */
	appointmentBook2.Patients = nil
	out, err = option.Marshal(&appointmentBook2) 
	if err != nil {
		fmt.Println(err)
	}

	var appointmentBook3 pb.AppointmentBook

	option_ := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	
	// err = protojson.Unmarshal(out, &appointmentBook3)
	// fmt.Println(appointmentBook3)
	err = option_.Unmarshal(out, &appointmentBook3)
	fmt.Println(appointmentBook3)
	
}

func getPerson() *pb.Person {
	return &pb.Person{
		Name:      "Dr. John Doe",
		Age:       31,
		Height:    1.75,
		Weight:    75.0,
		IsMarried: false,
	}
}

func getAppointmentBook() *pb.AppointmentBook {
	return &pb.AppointmentBook{
		Name: "Dr. John Doe",
		AppointmentDate: &pb.AppointmentDate{
			Day: pb.Day_SUNDAY,
			Date: 17,
			Month: pb.Month_APRIL,
			Year: 2022,
		},
		Patients: []*pb.Person{
			{
				Name:      "James Doe",
				Age:       31,
				Height:    1.75,
				Weight:    75.0,
				IsMarried: false,
			},
			{
				Name:      "Jenny Doe",
				Age:       32,
				Height:    1.65,
				Weight:    65.0,
				IsMarried: true,
			},
		},
		ContactDetails: map[string]string{
			"email": "john.doe@gmail.com",
			"phone": "1234567890",
		},
		// AppointmentStatus: &pb.AppointmentBook_IsCancelled{false},
		AppointmentStatus: &pb.AppointmentBook_Reason{"Appointment cancelled"},

	}
}
