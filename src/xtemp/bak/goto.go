package main

import "fmt"

func main() {
	a := 10
	if a > 1 {
		goto END
	}
	fmt.Println("222")
	err := 3
	err := 5
END:
	fmt.Println("END")
}
