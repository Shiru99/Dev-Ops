package main

import (
	"fmt"
	"net/http"
)

// Main Routine - Main Routine is created when we launch the program
func main(){
	links := []string{
		"https://www.google.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.google.kom",
	}

	for _,link := range links{
		// Child Go Routines - created with go keyword (similar to threads)
		go checkWebsiteStatus(link)
	}
}

func checkWebsiteStatus(link string){
	_,err := http.Get(link)
	if err != nil{
		fmt.Println(link, "is down!")
		return
	}
	fmt.Println(link, "is up!")
}