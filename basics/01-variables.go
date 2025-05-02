package main

import (
	"fmt"
	"math"
)

func main() {
	// Showing some values
	fmt.Println("Hello in Golang !")
	fmt.Println("Go" + "lang")
	fmt.Println("3+2=", 3+2)
	fmt.Println("7.0/4.0 = ", 7.0/4.0)
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!true)

	// Variables
	var name = "Anir"
	var age, username = 19, "4hbane"
	var isNice = true
	var int_example int // -> 0
	city := "Boston"    // Allowed only inside functions; short form for  ->  city string = "Boston"

	// %v: any  |  %d: integer  |  %s: string  |  %t: boolean
	fmt.Printf("name: %s, age: %d, username: %s, isNice: %t,  int_example: %d, fruit: %s\n", name, age, username, isNice, int_example, city)

	// Constants
	const pi = 3.14
	const result = 1e10 / pi
	fmt.Println(float32(result))
	fmt.Println(float32(math.Sin(result)))
}
