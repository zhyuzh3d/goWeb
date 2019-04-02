package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	a := time.Now().Unix()
	n, _ := strconv.Atoi(a)
	fmt.Println(">>>", n-1000)
}
