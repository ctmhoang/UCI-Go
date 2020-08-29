package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	str = strings.ToLower(str)


	if strings.HasPrefix(str, "i") &&
		strings.HasSuffix(str, "n") &&
		strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
