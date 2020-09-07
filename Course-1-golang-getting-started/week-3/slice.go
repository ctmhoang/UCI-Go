package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi! Please enter integers one after another and let me sort that for ya")
	fmt.Println("If you want to exit press 'X'")
	var numSlice = make([]int, 0, 3)
	var scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a number: ")
		scanner.Scan()
		in := scanner.Text()
		if i, err := strconv.Atoi(in); err == nil {
			numSlice = append(numSlice, i)
		} else if strings.EqualFold(in, "X") {
			fmt.Println("Byeee")
			break
		} else {
			fmt.Println("Doesnt seem like a number. Please re-enter again!")
			continue
		}
		sort.Ints(numSlice)
		fmt.Println("RESULT: ")
		fmt.Println(numSlice)
	}
}
