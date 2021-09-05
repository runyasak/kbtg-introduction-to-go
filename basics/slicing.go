package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "javascript"}
	fmt.Printf("langs: %#v\n", langs)

	a := langs[0:2] // [0, 2)
	fmt.Printf("a     : %#v\n", a)
	fmt.Printf("langs : %#v\n", langs)

	b := langs[1:3] // [1, 3)
	fmt.Printf("b     : %#v\n", b)
	fmt.Printf("langs : %#v\n", langs)

	n := langs[0:]
	fmt.Printf("n     : %#v\n", n)
	fmt.Printf("langs : %#v\n", langs)

	printSlice(langs)
	cords := [4]string{"Am", "B", "G", "F#"}
	printSlice(cords[:])
}

func printSlice(ns []string) {
	fmt.Printf("printSlice : ns: %#v\n", ns)
}
