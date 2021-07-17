package main

import "fmt"

func info(name, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
}

func today() string {
	return "มื้อนี่"
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	info("Gopher", "gopher เด้อ", 11)
	fmt.Println()
	info("Nong", "ซ่ำบายดี", 15)
	fmt.Println()

	day := today()
	fmt.Println(day)

	a, b := swap("ซ่ำบายดี", "Gopher")
	fmt.Println(a, b)
}
