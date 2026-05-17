package main

// HTTP GET request
// we need to output the HTTP response "code" and the body
import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {

	args := os.Args
	fmt.Println("\n Welcome to the HTTP client app \n")
	if len(args) < 2 {
		// case: no url is passed to CLI
		fmt.Println("\n usage: go run main.go <url>\n")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	// validate URL
	// parsing the URL => analyze it
	/*
		   split the URL into compomemts
			https://www.google.com/users
		   Validate the basic formartting: by returning 2 values
			1. parsedURL ==  *url.URL : struct ptr
				Scheme => https/http
				Host => google.com
				Path => /users
			2. err == validation for the state of the request
				err == nil "success"
				err != nil "failed"
	*/
	rawURL := os.Args[1]
	// we don't need to validate the url itself so we might escape it
	_, err := url.ParseRequestURI(rawURL)
	// parsedURL, err := url.ParsedRequestURI(rawURL)
	if err != nil {
		fmt.Println("\n Invalid URL %s \n", err)
		fmt.Println("Disconecting...\n")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	// http request, will go check if exists or not
	/*
		send HTTP request conects to the server and test the network reachability
		so checking:
			1. DNS resolution
			2. internet connectivity
			3. server availabilty
			*4. HTTP response
	*/
	// it returns 2 values, the response & err
	// the response is struct to *http.Response
	// status code, headers, body, content length,...
	// resp.Status, resp.StatusCode, resp.Header, resp.Body
	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Println("Request failed:\n", err)
		log.Fatal(err)

	} else {
		fmt.Println("The status code is: %s \n", resp.Status)
		fmt.Println("The URL is:\n", resp.Request.URL.String())
	}
	// the body is stream, and might be bigger than memory
	// it's stream, and we need to close that stream once we finish
	// defer => make sure that this runs at the end of the function
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
	log.Fatal(err)
	}
}
