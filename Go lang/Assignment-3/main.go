package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename:= os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/************** M-1 ***************/

	// io.Copy(os.Stdout, file)  /* Prints the response body to stdout */

	/************** M-2 : Custom Writer ***************/

	customWriter := myWriter{}
	io.Copy(customWriter, file)  // Writer interface (Destination), Reader interface (Source)

}

type myWriter struct{}

func (myWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
