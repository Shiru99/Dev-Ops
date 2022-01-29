package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main()  {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting2(sb)
}

// func (eb englishBot) getGreetings() string{
func (englishBot) getGreetings() string{
	return "Hi There!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreetings())
}


func (spanishBot) getGreetings() string{
	return "Hola!"
}

func printGreeting2(sb spanishBot) {
	fmt.Println(sb.getGreetings())
}