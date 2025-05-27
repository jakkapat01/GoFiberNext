package greeting

import (
	"fmt"

	personal "github.com/jakkapat01/GoFiberNext/greeting/internal"
)

// public function
func SayGreeting() {
	fmt.Println("Hello, from the greeting method")
	sayhi()
	personal.PrintPersonal()
}
