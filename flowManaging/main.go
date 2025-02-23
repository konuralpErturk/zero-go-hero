package main

import (
	"fmt"
	"time"
)

func main() {
	//lessonIfStatement()
	//lessonSwitchCase()
}

func lessonSwitchCase() {
	Weekday := time.Now().UTC().Weekday()
	dayOfWeek := int(Weekday)

	switch dayOfWeek {
	case 1:
		fmt.Println("It's a monday")
	case 2:
		fmt.Println("It's a thuesday")
	case 3:
		fmt.Println("It's a wednesday")
	case 4:
		fmt.Println("It's a thursday")
	case 5:
		fmt.Println("It's a friday")
	default:
		fmt.Println("It's the weekend")

	}
}

func lessonIfStatement() {
	value := -21
	var result string

	if value < 0 {
		result = "value lte zero"
	} else if value == 0 {
		result = "value equals zero"
	} else {
		result = "value gte zero"
	}

	fmt.Println(result)
}
