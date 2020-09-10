package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type name struct {
		fname string
		lname string
	}
	fmt.Println("File name? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	file, err := os.Open(scanner.Text())
	if err != nil{
		fmt.Fprintf(os.Stderr,"There are no such file like the name you provide")
		os.Exit(1)
	}
	defer file.Close()

	var data = make([]name,0,10)

	fscanner := bufio.NewScanner(file)
	for fscanner.Scan(){
		line := fscanner.Text();
		info := strings.Split(line," ")
		data = append(data, name{
			fname: info[0],
			lname: info[1],
		})
	}

	for _,val := range data{
		fmt.Println(val)
	}
}
