package function

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// set routing
func route() {
	functions.HTTP("HelloWorld", helloWorld)
	functions.HTTP("fizzBuzz", fizzBuzz)
}
