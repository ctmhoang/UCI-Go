package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num float64
	fmt.Println("Enter a floating point number and I will truncate it!")
	_, _ = fmt.Scan(&num)
	convted := strconv.FormatFloat(num,'f',0,32)
	fmt.Println(convted)
}