package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter your address: ")
	scanner.Scan()
	addr := scanner.Text()

	var infos  = map[string]string {"name" : name, "address" : addr}
	barray, _ :=json.Marshal(infos)
	fmt.Printf("%s",barray)
}
