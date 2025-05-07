package main

import "fmt"

func main() {

	fmt.Println("Basic loop : ")
	// Basic Loop with single condition
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("\nClassic loop :")
	// Classic initial / condition / for loop
	for j := 3; j <= 7; j++ {
		fmt.Println(j)
	}

	fmt.Println("\nLoop with range kw :")
	// do this n times with range - starts from 0
	for v := range 2 {
		fmt.Println("range", v)
	}

	fmt.Println("\nLoop without any condition :")
	// without any condition
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("\nUse break and continue kws:")
	// Availibility of break and continue in Go
	for v := range 10 {
		if v == 0 {
			fmt.Printf("Skip %d \n", v)
			continue // Go to the next iteration
		} else if v%2 == 0 {
			fmt.Printf("%d is even \n", v)
			break // stop the loop
		}
		fmt.Printf("%d is odd \n", v)
	}

	fmt.Println("\nIterate Slice of strings:")
	// Iterate Slice and use/ignioe the iteration index
	nicknames := []string{"Anir", "Steve", "Ayour"}
	for _, name := range nicknames { // ignore the index with _
		fmt.Println(name)
	}
	for k, name := range nicknames { // use the index with k
		fmt.Printf("Index %d: %s \n", k, name)
	}
}
