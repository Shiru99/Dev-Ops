package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func heyGoogle() {
	fmt.Println("from heyGoogle")
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Printf("Error : %v", error)
		os.Exit(1)
	}

	fmt.Println("******************************************************")
	fmt.Printf("%+v\n\n", resp)

	fmt.Printf("%T\n", resp.Body) // *http.gzipReader
	/* Interface :

		// Takes empty byte slice and put the response body into it
		type Reader interface {
			Read(empty []byte) (n int, err error)
		}

		// Takes data byte slice and writes it to cooresponding writer (terminal, file, etc)
		type Writer interface {
			Write(data []byte) (n int, err error)
		}
	*/

	/*

		// gzipReader wraps a response body so it can lazily
		// call gzip.NewReader on the first call to Read
		type gzipReader struct {
			_    incomparable
			body *bodyEOFSignal // underlying HTTP/1 response body framing
			zr   *gzip.Reader   // lazily-initialized gzip reader
			zerr error          // any error from gzip.NewReader; sticky
		}

		func (gz *gzipReader) Read(p []byte) (n int, err error) {
			if gz.zr == nil {
				if gz.zerr == nil {
					gz.zr, gz.zerr = gzip.NewReader(gz.body)
				}
				if gz.zerr != nil {
					return 0, gz.zerr
				}
			}

			gz.body.mu.Lock()
			if gz.body.closed {
				err = errReadOnClosedResBody
			}
			gz.body.mu.Unlock()

			if err != nil {
				return 0, err
			}
			return gz.zr.Read(p)
		}

		func (gz *gzipReader) Close() error {
			return gz.body.Close()
		}

	*/
	fmt.Printf("%+v\n", resp.Body)

	fmt.Println("******************************************************")

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// OR
	
	// io.Copy(os.Stdout, resp.Body) // Writer interface (Destination), Reader interface (Source)

	fmt.Println("*************************** Custom Interface ***************************")

	// Custom Interface

	customWriter := myWriter{}
	io.Copy(customWriter, resp.Body)

	fmt.Println("******************************************************")


}

type myWriter struct{}

// Implement Writer interface
func (myWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("My Writer has written this many bytes: ", len(bs))
	return len(bs), nil
}