package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "javascript"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("langs[0]: %#v\n", n)

	langs[1] = "Pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	l := len(langs)
	fmt.Printf("length of langs : %#v", l)

	langs = append(langs, "F#")
	fmt.Printf("langs: %#v\n", langs)

	fmt.Println("length of langs: ", len(langs))
}
