package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	langs := []string{"golang", "python", "javascript", "F#"}
	fmt.Println("\nfor basic")

	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Println(i, ":", value)
	}

	fmt.Println("\nfor range slice")
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}

	fmt.Println("\nonly value")
	for _, value := range langs {
		fmt.Println("only value:", value)
	}

	fmt.Println("\nonly index")
	for index := range langs {
		fmt.Println("only index:", index)
	}
}
