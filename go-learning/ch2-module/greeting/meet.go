package greeting

import (
	"fmt"
)

// public function
func SayMeeting() {
	fmt.Println("Hello, from the meeting method")
}

// private function
func sayhi() {
	fmt.Println("Hello, from the private method")
}
