package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		400: "Bad Reuest",
		500: "Internal Server Error",
	}
	fmt.Printf("% #v\n", status)

	l := len(status)
	fmt.Printf("length: %d\n\n", l)

	status[200] = "Okie"
	status[285] = "เมาแน่เลย"
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	value := status[200]
	fmt.Printf("%#v\n", value)

	delete(status, 285)
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	if v, ok := status[66]; ok {
		fmt.Printf("value: %q\n\n %v", v, ok)
	} else {
		fmt.Println("not found")
	}

	var m = make(map[string]string)

	if m == nil {
		fmt.Printf("m is nil. value: %#v\n", m)
	} else {
		fmt.Printf("m is not nil. value: %#v\n", m)
	}
}