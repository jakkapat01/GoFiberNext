package main

import (
	"fmt"
)

func IfCondition() {
	num := 15
	if num > 0 {
		fmt.Printf("%d is Positive number", num)
	} else {
		fmt.Printf("%d is Negative number", num)
	}

}
func IfElseIfCondition() {
	number1 := 10
	number2 := 20
	if number1 > number2 {
		fmt.Printf("%d is greater than %d", number1, number2)
	} else if number1 < number2 {
		fmt.Printf("%d is less than %d", number1, number2)
	} else {
		fmt.Printf("%d is equal to %d", number1, number2)
	}
}
func SwitchCase() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

func SwitchCaseMultipleCases() {
	dayOfWeek := "tuesday"
	switch dayOfWeek {
	case "sunday", "saturday":
		fmt.Println("It's a weekend")
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("It's a weekday")
	default:
		fmt.Println("Invalid day")
	}
}
