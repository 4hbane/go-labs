package main

import (
	"fmt"
	"time"
)

func main() {

	// IF - Else
	if 4%2 == 0 && 8%2 == 0 {
		fmt.Println("4 and 9 are even")
	}

	if age := 19; age < 18 {
		fmt.Println("Minor")
	} else if age < 65 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	// Switch for multiple if-else statments
	// Switch without an expression
	month := "March"
	switch {
	case month == "January": // no need for break
		fmt.Println("Start of the year")
	case month == "December":
		fmt.Println("End of the year")
	default:
		fmt.Println("Some other month")
	}

	// Switch with an expression
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Switch without an expression
	getType := func(variable interface{}) {
		switch t := variable.(type) {
		case bool:
			fmt.Println("This is a Boolean")
		case int:
			fmt.Println("This is a Integer")
		default:
			fmt.Printf("This is %T\n", t)
		}
	}

	getType(10)
	getType(false)
	getType([]string{"Black", "Yellow"})
}
