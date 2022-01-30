package main

import "fmt"

// Interface type : must implement all functions
type bot interface {
	getGreetings() string
	// getBotVersion() float64 	
}

type englishBot struct{}
type spanishBot struct{}

func main()  {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	heyGoogle()
}

// func (eb englishBot) getGreetings() string{ 		// Receiver function
func (englishBot) getGreetings() string{
	return "Hi There!"
}

func (spanishBot) getGreetings() string{
	return "Hola!"
}

func printGreeting(b bot){
	fmt.Println(b.getGreetings())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreetings())
// }

// func printGreeting2(sb spanishBot) {
// 	fmt.Println(sb.getGreetings())
// }