package main

import (
	"fmt"

	"github.com/runyasak/kbtg-introduction-to-go/sample-project/time"
	"github.com/runyasak/kbtg-introduction-to-go/sample-project/user"
)

func main() {
	user.Info("Nong", "โกเฟอร์", 15)

	t := time.Today()
	fmt.Println("today is ", t)
}
