package main

import (
	"fmt"
	"net/http"
	"time"
)

// Main Routine - Main Routine is created when we launch the program
func main() {
	links := []string{
		"https://www.google.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		// "https://www.google.kom",
	}

	c := make(chan string)

	for _, link := range links {
		// Child Go Routines - created with go keyword (similar to threads)
		go checkWebsiteStatus(link, c)
	}

	/*
		Channel pauses all other Routines (including Main Routine) until the data is received from the channel
	*/
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	/* infinite loop */
	// for {
	// 	checkWebsiteStatus(<-c,c)
	// }

	for l := range c {
		// go checkWebsiteStatus(l, c)
		/* Literal Function (lambda function) */
		go func(link string)  {
			time.Sleep(5 * time.Second)
			checkWebsiteStatus(link, c)
		}(l)
	}
}

func checkWebsiteStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- link // Send the string data to the channel
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
