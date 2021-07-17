package main

import "fmt"

var name = "Nong"

var say = func(n string) {
	fmt.Printf("May name is %s\n", n)
}

var cal = func(op func(int, int) int) {
	r := op(4, 5)
	fmt.Printf("result : %v\n", r)
}

func main() {
	fmt.Printf("value: %v\n", name)
	fmt.Printf("type: %T\n", name)

	say(name)

	plus := func(a, b int) int {
		return a + b
	}

	p := plus(1, 2)
	fmt.Println("1+2 =", p)
	fmt.Printf("type: %T\n", plus)

	cal(plus)

	minus := func(a, b int) int { return a - b }
	cal(minus)
}
