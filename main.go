package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n') // Read until newline
	city = city[:len(city)-1]          // Remove trailing newline
	fmt.Println("You live in", city)
}
