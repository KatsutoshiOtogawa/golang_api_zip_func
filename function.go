package function

import (
	"fmt"
	"net/http"
)

// init function. this function call.
func init() {
    route()
}

// helloWorld writes "Hello, World!" to the HTTP response.
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// helloWorld writes "Hello, World!" to the HTTP response.
func fizzBuzz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "fizzBuzz")
}
