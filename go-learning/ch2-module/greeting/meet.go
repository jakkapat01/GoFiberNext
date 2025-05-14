package greeting

import (
	"fmt"
)

// public function
func SayMeeting() {
	fmt.Println("Meet, from greet package")
}

// private function
func sayhi() {
	fmt.Println("hi private, from greet package")
}
